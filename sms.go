package qcloudsms

import (
	"io"

	"io/ioutil"

	"encoding/json"

	"github.com/pkg/errors"
)

type SmsExt struct {
	Type   int
	Extend string
	Ext    string
	Max    int
}

type SMS struct {
	client Client
}

var defaultExt = SmsExt{Max: 10}

func SMSService(client Client) (*SMS, error) {
	if client == nil {
		return nil, errors.New("client not init")
	}
	return &SMS{
		client: client,
	}, nil
}

// Tel to store mobile phone number.
type Tel struct {
	NationCode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
}

// Send sms to one user. phone like `18600000000`.
func (ss *SMS) Send(phone, content string, exts ...SmsExt) (resp *QResponse, err error) {
	var ext SmsExt
	if len(exts) > 0 {
		ext = exts[0]
	} else {
		ext = defaultExt
	}
	req := &QRequest{
		Path:   "sendsms",
		Type:   ext.Type,
		Msg:    content,
		Extend: ext.Extend,
		Ext:    ext.Ext,
	}
	req.Tel = Tel{
		// Dirty work for Chinese only telephone.
		NationCode: "86",
		Mobile:     phone,
	}
	resp, err = ss.client.Post(req)
	return
}

func (ss *SMS) MultiSend(phones []string, content string, exts ...SmsExt) (resp *QResponse, err error) {
	var ext SmsExt
	if len(exts) > 0 {
		ext = exts[0]
	} else {
		ext = defaultExt
	}
	if len(phones) == 0 {
		err = errors.New("No phone number given")
		return
	}
	req := &QRequest{
		Path:   "sendmultisms2",
		Type:   ext.Type,
		Msg:    content,
		Extend: ext.Extend,
		Ext:    ext.Ext,
	}
	tel := make([]Tel, len(phones))
	for i, p := range phones {
		tel[i] = Tel{
			// Dirty work for Chinese only telephone.
			NationCode: "86",
			Mobile:     p,
		}
	}
	req.Tel = tel
	resp, err = ss.client.Post(req)
	return
}

func (ss *SMS) MobileStatus(mobile string, start, end int64, exts ...SmsExt) (resp *QResponse, err error) {
	var ext SmsExt
	if len(exts) > 0 {
		ext = exts[0]
	} else {
		ext = defaultExt
	}
	if ext.Max > 100 {
		ext.Max = 100
	}
	req := &QRequest{
		Path:       "pullstatus4mobile",
		Type:       ext.Type,
		Max:        ext.Max,
		BeginTime:  start,
		EndTime:    end,
		NationCode: "86",
		Mobile:     mobile,
	}
	resp, err = ss.client.Post(req)
	return
}

func (ss *SMS) ReadStatus(reader io.Reader) (status []QSMSStatus, err error) {
	var b []byte
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	status, err = ss.ReadStatusByte(b)
	return
}

func (ss *SMS) ReadStatusByte(body []byte) ([]QSMSStatus, error) {
	var status = make([]QSMSStatus, 0)
	err := json.Unmarshal(body, &status)
	return status, err
}

func (ss *SMS) ReadReply(reader io.Reader) (status QSMSReply, err error) {
	var b []byte
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	status, err = ss.ReadReplyByte(b)
	return
}

func (ss *SMS) ReadReplyByte(body []byte) (QSMSReply, error) {
	var status QSMSReply
	err := json.Unmarshal(body, &status)
	return status, err
}
