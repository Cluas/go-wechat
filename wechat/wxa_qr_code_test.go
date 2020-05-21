package wechat

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestWXAService_GetWXACode_Stream(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wxa/getwxacode", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=hello-world.jpg")
		fmt.Fprint(w, "Hello World")
	})

	resp, err := client.WXA.GetWXACode(context.Background(), "o", &GetWXACodeRequest{
		Path:      "",
		Width:     0,
		AutoColor: false,
		LineColor: nil,
		IsHyaline: false,
	})
	if err != nil {
		t.Errorf("WXA.GetWXACode returned error: %v", err)
	}
	want := []byte("Hello World")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("WXA.GetWXACode returned bad reader: %v", err)
	}
	if !bytes.Equal(want, content) {
		t.Errorf("WXA.GetWXACode returned %+v, want %+v", content, want)
	}
}

func TestWXAService_CreateWXAQRCode_Stream(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/cgi-bin/wxaapp/createwxaqrcode", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=hello-world.jpg")
		fmt.Fprint(w, "Hello World")
	})

	resp, err := client.WXA.CreateWXAQRCode(context.Background(), "o", &CreateQRCodeRequest{
		Path:  "",
		Width: 0,
	})
	if err != nil {
		t.Errorf("WXA.CreateWXAQRCode returned error: %v", err)
	}
	want := []byte("Hello World")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("WXA.CreateWXAQRCode returned bad reader: %v", err)
	}
	if !bytes.Equal(want, content) {
		t.Errorf("WXA.CreateWXAQRCode returned %+v, want %+v", content, want)
	}
}

func TestWXAService_GetWXACodeUnlimit_Stream(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wxa/getwxacodeunlimit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=hello-world.jpg")
		fmt.Fprint(w, "Hello World")
	})

	resp, err := client.WXA.GetWXACodeUnlimit(context.Background(), "o", &GetWXACodeUnlimitRequest{
		Scene:     "",
		Path:      "",
		Width:     0,
		AutoColor: false,
		LineColor: nil,
		IsHyaline: false,
	})
	if err != nil {
		t.Errorf("WXA.GetWXACodeUnlimit returned error: %v", err)
	}
	want := []byte("Hello World")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("WXA.GetWXACodeUnlimit returned bad reader: %v", err)
	}
	if !bytes.Equal(want, content) {
		t.Errorf("WXA.GetWXACodeUnlimit returned %+v, want %+v", content, want)
	}
}

func TestWXAService_GetQrCode_Stream(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/wxa/get_qrcode", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Disposition", "attachment; filename=hello-world.jpg")
		fmt.Fprint(w, "Hello World")
	})

	resp, err := client.WXA.GetQrCode(context.Background(), "o", "")
	if err != nil {
		t.Errorf("WXA.GetQrCode returned error: %v", err)
	}
	want := []byte("Hello World")
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("WXA.GetQrCode returned bad reader: %v", err)
	}
	if !bytes.Equal(want, content) {
		t.Errorf("WXA.GetQrCode returned %+v, want %+v", content, want)
	}
}
