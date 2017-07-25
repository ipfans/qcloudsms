package qcloudsms

import "github.com/pkg/errors"

type SmsExt struct {
	Type   int
	Extend string
	Ext    string
}

type SMS struct {
	client Client
}

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
		ext = SmsExt{}
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
