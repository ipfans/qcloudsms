package qcloudsms

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleVoiceMSSend() {
	conf := NewClientConfig()
	conf.AppID = "appID"
	conf.AppKey = "appKey"
	client, err := NewClient(conf)
	if err != nil {
		return
	}
	voice, err := VoiceMessageService(client)
	if err != nil {
		return
	}
	resp, err := voice.Send("18600000000", "1234")
	if err != nil {
		return
	}
	fmt.Println(resp.ErrMsg)
}

func TestVoiceMSSend(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		phone   string
		content string
		exts    []VoiceExt
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *QResponse
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vms := &VoiceMS{
				client: tt.fields.client,
			}
			gotResp, err := vms.Send(tt.args.phone, tt.args.content, tt.args.exts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("VoiceMS.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("VoiceMS.Send() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestVoiceMSSendPrompt(t *testing.T) {
	type fields struct {
		client Client
	}
	type args struct {
		phone   string
		content string
		exts    []VoiceExt
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *QResponse
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vms := &VoiceMS{
				client: tt.fields.client,
			}
			gotResp, err := vms.SendPrompt(tt.args.phone, tt.args.content, tt.args.exts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("VoiceMS.SendPrompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("VoiceMS.SendPrompt() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
