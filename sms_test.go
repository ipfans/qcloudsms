package qcloudsms

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

// var (
// 	appID       = ""
// 	appKey      = ""
// 	testPhone   = ""
// 	testContent = ""
// )

// func TestRealSend(t *testing.T) {
// 	conf := NewClientConfig()
// 	conf.AppID = appID
// 	conf.AppKey = appKey
// 	client, err := NewClient(conf)
// 	if err != nil {
// 		t.Errorf("Client init failed: %v", err)
// 		return
// 	}
// 	sms, err := SMSService(client)
// 	if err != nil {
// 		t.Errorf("Sms init failed: %v", err)
// 		return
// 	}
// 	resp, err := sms.Send(testPhone, testContent)
// 	if err != nil {
// 		t.Errorf("Send request failed: %v", err)
// 		return
// 	}
// 	if !cmp.Equal(resp.Error(), ErrCodeOK) {
// 		t.Errorf("Send sms failed: %v")
// 	}
// }

func TestSend(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	c.EXPECT().Post(gomock.Any()).Return(&QResponse{
		ErrMsg: "OK",
		Sid:    "xxxxxxx",
		Fee:    1,
	}, nil)
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
