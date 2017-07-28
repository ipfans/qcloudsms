package qcloudsms

import (
	"net"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	nonMobileSig   = "27616b6f91b2af339e2fbebefa8e7f095820519bb0c003b1e71f494c122f1c2d"
	mobileSig      = "d88472b109a599a4631160fac0f59cfc30fe6c59547b58d6c7225bd5cb9f3fc6"
	mutliMobileSig = "44edc2dfe23d4e4a89d92be04c23cdf5b5d69fa59470918c2830261126a24d19"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient(nil)
	if err == nil {
		t.Errorf("NewClient can init with empty config")
		return
	}
	_, err = NewClient(NewClientConfig())
	if err == nil {
		t.Errorf("NewClient can init with empty config")
		return
	}
}

func TestPost(t *testing.T) {
	oldBaseUrl := baseURL
	ln, _ := net.Listen("tcp4", "127.0.0.1:0")
	l := ln.(*net.TCPListener)
	srv := &http.Server{Addr: "127.0.0.1:8081"}
	http.HandleFunc("/v5/tlsvoicesvr/sendsms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"result":0,"errmsg":"OK","ext":"","sid":"123123123","fee":1}`))
	})
	go srv.Serve(l)
	defer srv.Shutdown(nil)
	baseURL = "http://" + l.Addr().String() + "/v5"
	conf := NewClientConfig()
	conf.AppID = "123"
	conf.AppKey = "dffdfd6029698a5fdf4"
	client, _ := NewClient(conf)
	sms, err := MessageService(client)
	if err != nil {
		t.Errorf("Sms init failed: %v", err)
		return
	}
	resp, err := sms.Send("18600000000", "Hello world")
	if err != nil {
		t.Errorf("Send request failed: %v", err)
		return
	}
	if !cmp.Equal(resp.Error(), ErrCodeOK) {
		t.Errorf("Send sms failed: %v", resp.Error())
		return
	}
	baseURL = oldBaseUrl
}

func TestSignature(t *testing.T) {
	conf := NewClientConfig()
	conf.AppID = "123"
	conf.AppKey = "dffdfd6029698a5fdf4"
	client, _ := NewClient(conf)
	sig := client.(*smsClient).signature(nil, "7226249334", 1457336869)
	if !cmp.Equal(sig, nonMobileSig) {
		t.Errorf("sig want be %s, but it was %s", nonMobileSig, sig)
		return
	}
	sig = client.(*smsClient).signature("", "7226249334", 1457336869)
	if !cmp.Equal(sig, nonMobileSig) {
		t.Errorf("sig want be %s, but it was %s", nonMobileSig, sig)
		return
	}
	sig = client.(*smsClient).signature(Tel{Mobile: "18600000000"}, "7226249334", 1457336869)
	if !cmp.Equal(sig, mobileSig) {
		t.Errorf("sig want be %s, but it was %s", mobileSig, sig)
		return
	}
	sig = client.(*smsClient).signature([]Tel{{Mobile: "18600000000"}, {Mobile: "15900000000"}}, "7226249334", 1457336869)
	if !cmp.Equal(sig, mutliMobileSig) {
		t.Errorf("sig want be %s, but it was %s", mutliMobileSig, sig)
	}
}
