package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_Plugin_Apply(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/plugin", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `{"errcode":0,"errmsg":"ok","plugin_list":[],"apply_list":[]}`)
	})
	got, _, err := client.WXA.Plugin(context.Background(), "token", &PluginRequest{
		Action:      "apply",
		PluginAppID: "",
		UserVersion: "",
	})
	if err != nil {
		t.Errorf("WXA.GetLiveInfo retured err: %v", err)
	}
	want := &PluginResponse{
		PluginList: []*Plugin{},
		ApplyList:  []*Plugin{},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("WXA.Plugin got %+v, want %+v", got, want)
	}

}

func TestWXAService_Plugin_List(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/plugin", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprintf(w, `
							{
							  "errcode": 0,
							  "errmsg": "ok",
							  "plugin_list": [
								{
								  "appid": "aaaa",
								  "status": 1,
								  "nickname": "插件昵称",
								  "headimgurl": "http://plugin.qq.com"
								}
							  ],
                               "apply_list":[]
							}
					`)
	})
	got, _, err := client.WXA.Plugin(context.Background(), "token", &PluginRequest{
		Action:      "list",
		PluginAppID: "",
		UserVersion: "",
	})
	if err != nil {
		t.Errorf("WXA.GetLiveInfo retured err: %v", err)
	}
	want := &PluginResponse{
		PluginList: []*Plugin{
			{
				AppID:      "aaaa",
				Status:     1,
				Nickname:   "插件昵称",
				HeadImgURL: "http://plugin.qq.com",
			},
		},
		ApplyList: []*Plugin{},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("WXA.Plugin got %+v, want %+v", got, want)
	}

}
