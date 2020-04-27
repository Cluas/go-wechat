package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_BindTester(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &BindTesterRequest{WechatID: "testid"}

	mux.HandleFunc("/wxa/bind_tester", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "userstr": "xxxxxxxxx"
							}`)
	})

	got, _, err := client.WXA.BindTester(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.Tester retured err: %v", err)
	}
	want := &Tester{UserString: "xxxxxxxxx"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.Tester got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetShowWXAItem(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/getshowwxaitem", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "can_open": 1,
							  "is_open": 1,
							  "appid": "展示的公众号appid",
							  "nickname": "展示的公众号nickname",
							  "headimg": "展示的公众号头像"
							}`)
	})
	got, _, err := client.WXA.GetShowWXAItem(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetShowWXAItem retured err: %v", err)
	}
	want := &ShowWXAItem{
		CanOpen:   1,
		IsOpen:    1,
		AppID:     "展示的公众号appid",
		Nickname:  "展示的公众号nickname",
		HeadImage: "展示的公众号头像",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetShowWXAItem got %+v, want %+v", got, want)
	}
}

func TestWXAService_MemberAuth(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/memberauth", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
								"errcode": 0,
								"errmsg": "ok",
								"members": [
									{
										"userstr": "xxxxxxxx"
									},
									{
										"userstr": "yyyyyyyy"
									}
								]
							}`)
	})
	got, _, err := client.WXA.MemberAuth(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.MemberAuth retured err: %v", err)
	}
	want := &Testers{Members: []*Tester{{UserString: "xxxxxxxx"}, {UserString: "yyyyyyyy"}}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.MemberAuth got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetWXAMpLinkForShow(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/getwxamplinkforshow", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "total_num": 10,
							  "biz_info_list": [
								{
								  "nickname": "公众号昵称",
								  "appid": "公众号appid",
								  "headimg": "公众号头像"
								}
							  ]
							}`)
	})

	got, _, err := client.WXA.GetWXAMpLinkForShow(context.Background(), "token", 1, 1)
	if err != nil {
		t.Errorf("WXA.GetWXAMpLinkForShow retured err: %v", err)
	}
	want := &WXAMPLinks{
		BIZInfoList: []*WXAMPLink{{Nickname: "公众号昵称", AppID: "公众号appid", HeadImage: "公众号头像"}},
		TotalNum:    10,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetWXAMpLinkForShow got %+v, want %+v", got, want)
	}
}

func TestWXAService_UnBindTester(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &UnBindTesterRequest{
		UserString: "testtest",
		WechatID:   "",
	}
	mux.HandleFunc("/wxa/unbind_tester", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})

	_, err := client.WXA.UnBindTester(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.UnBindTester retured err: %v", err)
	}

}

func TestWXAService_UpdateShowWXAItem(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &UpdateShowWXAItemRequest{
		WXASubscribeBIZFlag: 1,
		AppID:               "如果开启，需要传新的公众号appid",
	}

	mux.HandleFunc("/wxa/updateshowwxaitem", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})

	_, err := client.WXA.UpdateShowWXAItem(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.UpdateShowWXAItem retured err: %v", err)
	}
}
