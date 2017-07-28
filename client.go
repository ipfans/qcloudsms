package qcloudsms

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

var (
	userAgent   = "qcloudsms-go-sdk/1.0"
	baseURL     = "https://yun.tim.qq.com/v5"
	contentType = "application/json; charset=utf-8"
)

// ClientConfig to set configuration for client.
type ClientConfig struct {
	initialized bool

	AppID     string
	AppKey    string
	UserAgent string
}

// NewClientConfig returns instance for client config.
func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		initialized: true,

		UserAgent: userAgent,
	}
}

var defaultClientConfig = &ClientConfig{
	initialized: false,
	UserAgent:   userAgent,
}

// Client works for API Call.
type Client interface {
	Post(*QRequest) (*QResponse, error)
}

type smsClient struct {
	client *http.Client
	config *ClientConfig
}

// NewClient returns Client interface or error.
func NewClient(conf *ClientConfig) (Client, error) {
	if conf == nil {
		conf = defaultClientConfig
	}
	if conf.initialized == false {
		return nil, errors.New("Please init client configuration with NewClientConfig()")
	}
	if conf.AppID == "" || conf.AppKey == "" {
		return nil, errors.New("App information not configure")
	}
	return &smsClient{
		client: &http.Client{},
		config: conf,
	}, nil
}

// Post request to qcloud API.
func (sc *smsClient) Post(req *QRequest) (resp *QResponse, err error) {
	random := fmt.Sprintf("%06d", rand.Intn(999999))
	params := url.Values{}
	params.Set("sdkappid", sc.config.AppID)
	params.Set("random", random)
	targetURL := baseURL + req.Path + "?" + params.Encode()
	now := time.Now().Unix()
	req.Time = now
	req.Sig = sc.signature(req.Tel, random, now)
	b, e := json.Marshal(req)
	if e != nil {
		err = errors.Wrap(e, "Client request parse")
		return
	}
	httpReq, e := http.NewRequest("POST", targetURL, bytes.NewReader(b))
	if e != nil {
		err = errors.Wrap(e, "Client NewRequest")
		return
	}
	httpReq.Header.Set("User-Agent", sc.config.UserAgent)
	httpReq.Header.Set("Content-Type", contentType)
	httpResp, e := sc.client.Do(httpReq)
	if e != nil {
		err = errors.Wrap(e, "Client post")
		return
	}
	defer httpResp.Body.Close()
	b, e = ioutil.ReadAll(httpResp.Body)
	if e != nil {
		err = errors.Wrap(e, "Client read")
		return
	}
	resp = &QResponse{}
	e = json.Unmarshal(b, resp)
	if e != nil {
		err = errors.Wrap(e, "Client json parse")
		return
	}
	return
}

func (sc *smsClient) signature(tel interface{}, random string, now int64) string {
	var mobile string
	if tel != nil {
		switch tel.(type) {
		case Tel:
			mobile = tel.(Tel).Mobile
		case []Tel:
			for _, t := range tel.([]Tel) {
				if mobile == "" {
					mobile = t.Mobile
				} else {
					mobile += "," + t.Mobile
				}
			}
		default:
			mobile = ""
		}
	}
	s := fmt.Sprintf("appkey=%s&random=%s&time=%d", sc.config.AppKey, random, now)
	if mobile != "" {
		s += "&mobile=" + mobile
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}
