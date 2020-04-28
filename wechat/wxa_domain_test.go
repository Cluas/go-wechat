package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_ModifyDomain(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/modify_domain", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "requestdomain": ["https://www.qq.com", "https://www.qq.com"],
							  "wsrequestdomain": ["wss://www.qq.com", "wss://www.qq.com"],
							  "uploaddomain": ["https://www.qq.com", "https://www.qq.com"],
							  "downloaddomain": ["https://www.qq.com", "https://www.qq.com"]
							}`)
	})
	got, _, err := client.WXA.ModifyDomain(context.Background(), "token", &ModifyDomainRequest{
		Action:          "add",
		RequestDomain:   []string{"https://www.qq.com", "https://www.qq.com"},
		WSRequestDomain: []string{"wss://www.qq.com", "wss://www.qq.com"},
		UploadDomain:    []string{"https://www.qq.com", "https://www.qq.com"},
		DownloadDomain:  []string{"https://www.qq.com", "https://www.qq.com"},
	})
	if err != nil {
		t.Errorf("WXA.ModifyDomain retured err: %v", err)
	}
	want := &Domain{
		RequestDomain:   []string{"https://www.qq.com", "https://www.qq.com"},
		WSRequestDomain: []string{"wss://www.qq.com", "wss://www.qq.com"},
		UploadDomain:    []string{"https://www.qq.com", "https://www.qq.com"},
		DownloadDomain:  []string{"https://www.qq.com", "https://www.qq.com"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("WXA.ModifyDomain got %+v, want %+v", got, want)
	}
}

func TestWXAService_SetWebViewDomain(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/setwebviewdomain", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							}`)
	})
	_, err := client.WXA.SetWebViewDomain(context.Background(), "token", &SetWebViewDomainRequest{
		Action:        "add",
		WebViewDomain: []string{"https://www.qq.com", "https://m.qq.com"},
	})
	if err != nil {
		t.Errorf("WXA.SetWebViewDomain retured err: %v", err)
	}
}
