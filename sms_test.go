package qcloudsms

import (
	"fmt"
	"io"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
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
		Data: []QSMSStatus{
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
	resp, err := sms.MobileStatus("18600000000", 0, 1500000000)
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
		Data: []QSMSReply{
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
	resp, err = sms.MobileStatus("18600000000", 0, 1500000000)
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

func TestSMSReadStatus(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		reader io.Reader
	}
	data := `[{
        "user_receive_time": "2015-10-17 08:03:04",
        "nationcode": "86",
        "mobile": "13xxxxxxxxx",
        "report_status": "SUCCESS",
        "errmsg": "DELIVRD",
        "description": "用户短信送达成功",
        "sid": "xxxxxxx"
    }]`
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus []QSMSStatus
		wantErr    bool
	}{
		{
			name: "valid io.reader",
			fields: fields{
				c,
			},
			args: args{
				strings.NewReader(data),
			},
			wantStatus: []QSMSStatus{
				{
					UserReceiveTime: "2015-10-17 08:03:04",
					NationCode:      "86",
					Mobile:          "13xxxxxxxxx",
					ReportStatus:    "SUCCESS",
					ErrMsg:          "DELIVRD",
					Description:     "用户短信送达成功",
					Sid:             "xxxxxxx"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SMS{
				client: tt.fields.client,
			}
			gotStatus, err := ss.ReadStatus(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("SMS.ReadStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotStatus, tt.wantStatus) {
				t.Errorf("SMS.ReadStatus() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestSMSReadReply(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		reader io.Reader
	}
	data := `{
		"nationcode": "86",
		"mobile": "13xxxxxxxxx",
		"text": "用户回复的内容",
		"sign": "短信签名",
		"time": 1457336869,
		"extend": "扩展码"
	}`
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	c := NewMockClient(mockCtrl)
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus QSMSReply
		wantErr    bool
	}{
		{
			name: "valid io.Reader",
			fields: fields{
				c,
			},
			args: args{
				strings.NewReader(data),
			},
			wantStatus: QSMSReply{
				NationCode: "86",
				Mobile:     "13xxxxxxxxx",
				Text:       "用户回复的内容",
				Sign:       "短信签名",
				Time:       1457336869,
				Extend:     "扩展码",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SMS{
				client: tt.fields.client,
			}
			gotStatus, err := ss.ReadReply(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("SMS.ReadReply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotStatus, tt.wantStatus) {
				t.Errorf("SMS.ReadReply() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
