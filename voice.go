package qcloudsms

import "github.com/pkg/errors"

type VoiceExt struct {
	Ext        string
	PlayTimes  int
	PromptType int
}

type VoiceMS struct {
	client Client
}

var defaultVoiceExt = VoiceExt{
	PlayTimes:  3,
	PromptType: 2,
}

func VoiceMessageService(client Client) (*VoiceMS, error) {
	if client == nil {
		return nil, errors.New("client not init")
	}
	return &VoiceMS{
		client: client,
	}, nil
}

// Send voice message to one user. phone like `18600000000`.
func (vms *VoiceMS) Send(phone, content string, exts ...VoiceExt) (resp *QResponse, err error) {
	var ext VoiceExt
	if len(exts) > 0 {
		ext = exts[0]
	} else {
		ext = defaultVoiceExt
	}
	req := &QRequest{
		Path:      "/tlsvoicesvr/sendvoice",
		Msg:       content,
		Ext:       ext.Ext,
		PlayTimes: ext.PlayTimes,
	}
	req.Tel = Tel{
		// Dirty work for Chinese only telephone.
		NationCode: "86",
		Mobile:     phone,
	}
	resp, err = vms.client.Post(req)
	return
}

// SendPrompt to send voice message to one user. phone like `18600000000`.
func (vms *VoiceMS) SendPrompt(phone, content string, exts ...VoiceExt) (resp *QResponse, err error) {
	var ext VoiceExt
	if len(exts) > 0 {
		ext = exts[0]
	} else {
		ext = defaultVoiceExt
	}
	req := &QRequest{
		Path:       "/tlsvoicesvr/sendvoiceprompt",
		PromptFile: content,
		Ext:        ext.Ext,
		PlayTimes:  ext.PlayTimes,
		PromptType: ext.PromptType,
	}
	req.Tel = Tel{
		// Dirty work for Chinese only telephone.
		NationCode: "86",
		Mobile:     phone,
	}
	resp, err = vms.client.Post(req)
	return
}
