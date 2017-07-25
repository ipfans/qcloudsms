package qcloudsms

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func ExampleSend() {
	conf := NewClientConfig()
	conf.AppID = "appID"
	conf.AppKey = "appKey"
	client, err := NewClient(conf)
	if err != nil {
		return
	}
	sms, err := SMSService(client)
	if err != nil {
		return
	}
	resp, err := sms.Send("18600000000", "Hello World!")
	if err != nil {
		return
	}
	fmt.Println(resp.ErrMsg)
}

func TestSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	r := &QResponse{}
	r.ErrMsg = "OK"
	r.Sid = "xxxx"
	r.Fee = 1
	c.EXPECT().Post(gomock.Any()).Return(r, nil)
	sms, _ := SMSService(c)
	resp, err := sms.Send("18600000000", "Hello world")
	if err != nil {
		t.Errorf("Send request failed: %v", err)
		return
	}
	if resp.Error() != ErrCodeOK {
		t.Errorf("Send sms failed: %v", resp.Error())
		return
	}
	if resp.Desc() != "OK" {
		t.Errorf("Send desc failed: %v", resp.Desc())
		return
	}
}

func TestMultiSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	r := &QResponse{
		Result: 0,
		ErrMsg: "OK",
		Ext:    "",
		Detail: []QStatus{
			{
				Result:     0,
				ErrMsg:     "OK",
				Mobile:     "18600000000",
				NationCode: "86",
				Sid:        "xxxxxxx",
				Fee:        1,
			},
			{
				Result:     0,
				ErrMsg:     "OK",
				Mobile:     "15900000000",
				NationCode: "86",
				Sid:        "xxxxxxx",
				Fee:        1,
			},
		},
	}
	c.EXPECT().Post(gomock.Any()).Return(r, nil)
	sms, _ := SMSService(c)
	resp, err := sms.MultiSend([]string{"18600000000", "15900000000"}, "Hello world")
	if err != nil {
		t.Errorf("Send request failed: %v", err)
		return
	}
	if resp.Error() != ErrCodeOK {
		t.Errorf("Send sms failed: %v", resp.Error())
		return
	}
}

func TestMobileStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	r := &QResponse{
		Result: 0,
		ErrMsg: "OK",
		Ext:    "",
		Data: []DeliveryStatus{
			{
				UserReceiveTime: "2017-01-01 00:00:00",
				NationCode:      "86",
				Mobile:          "18600000000",
				ReportStatus:    "SUCCESS",
				ErrMsg:          "DELIVRD",
				Description:     "用户短信接收成功",
				Sid:             "xxxx",
			},
		},
	}
	c.EXPECT().Post(gomock.Any()).Return(r, nil)
	sms, _ := SMSService(c)
	resp, err := sms.MultiSend([]string{"18600000000", "15900000000"}, "Hello world")
	if err != nil {
		t.Errorf("Send request failed: %v", err)
		return
	}
	if resp.Error() != ErrCodeOK {
		t.Errorf("Send sms failed: %v", resp.Error())
		return
	}
	if len(resp.Reply()) != 0 {
		t.Errorf("Status parse failed: %v", resp.Error())
		return
	}
	if len(resp.Status()) == 0 {
		t.Errorf("Status parse failed: %v", resp.Error())
		return
	}
	c = NewMockClient(mockCtrl)
	r = &QResponse{
		Result: 0,
		ErrMsg: "OK",
		Ext:    "",
		Data: []ReplyStatus{
			{
				Time:       1500000000,
				NationCode: "86",
				Mobile:     "18600000000",
				Sign:       "【你好科技】",
				Text:       "HelloWorld",
				Extend:     "",
			},
		},
	}
	c.EXPECT().Post(gomock.Any()).Return(r, nil)
	sms, _ = SMSService(c)
	resp, err = sms.MultiSend([]string{"18600000000", "15900000000"}, "Hello world")
	if err != nil {
		t.Errorf("Send request failed: %v", err)
		return
	}
	if resp.Error() != ErrCodeOK {
		t.Errorf("Send sms failed: %v", resp.Error())
		return
	}
	if len(resp.Status()) != 0 {
		t.Errorf("Reply parse failed: %v", resp.Error())
		return
	}
	if len(resp.Reply()) == 0 {
		t.Errorf("Reply parse failed: %v", resp.Error())
		return
	}
	if resp.Reply()[0].Text != "HelloWorld" {
		t.Errorf("Content want be %s, but it was %s",
			"HelloWorld", resp.Reply()[0].Text)
		return
	}
}
