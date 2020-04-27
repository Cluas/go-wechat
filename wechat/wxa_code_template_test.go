package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_AddToTemplate(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/addtotemplate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							}`)
	})
	_, err := client.WXA.AddToTemplate(context.Background(), "token", 0)
	if err != nil {
		t.Errorf("WXA.AddToTemplate retured err: %v", err)
	}
}
func TestWXAService_DeleteTemplate(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/deletetemplate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "status": 1,
							  "reason": "帐号信息不合规范",
							  "screenshot": "xxx|yyy|zzz"
							}`)
	})
	_, err := client.WXA.DeleteTemplate(context.Background(), "token", 0)
	if err != nil {
		t.Errorf("WXA.DeleteTemplate retured err: %v", err)
	}

}

func TestWXAService_GetTemplateDraftList(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/gettemplatedraftlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "draft_list": [
								{
								  "create_time": 1488965944,
								  "user_version": "VVV",
								  "user_desc": "AAS",
								  "draft_id": 0
								},
								{
								  "create_time": 1504790906,
								  "user_version": "11",
								  "user_desc": "111111",
								  "draft_id": 4
								}
							  ]
							}`)
	})
	got, _, err := client.WXA.GetTemplateDraftList(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetTemplateDraftList retured err: %v", err)
	}
	want := &TemplateDrafts{DraftList: []*Draft{
		{
			CreateTime:      1488965944,
			UserVersion:     "VVV",
			UserDescription: "AAS",
			DraftID:         0,
		},
		{
			CreateTime:      1504790906,
			UserVersion:     "11",
			UserDescription: "111111",
			DraftID:         4,
		},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetTemplateDraftList got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetTemplateList(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/gettemplatelist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "template_list": [
								{
								  "create_time": 1488965944,
								  "user_version": "VVV",
								  "user_desc": "AAS",
								  "template_id": 0
								},
								{
								  "create_time": 1504790906,
								  "user_version": "11",
								  "user_desc": "111111",
								  "template_id": 4
								}
							  ]
							}`)
	})
	got, _, err := client.WXA.GetTemplateList(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetTemplateList retured err: %v", err)
	}
	want := &Templates{TemplateList: []*Template{
		{
			CreateTime:  1488965944,
			UserVersion: "VVV",
			UserDesc:    "AAS",
			TemplateID:  0,
		},
		{
			CreateTime:  1504790906,
			UserVersion: "11",
			UserDesc:    "111111",
			TemplateID:  4,
		},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetTemplateList got %+v, want %+v", got, want)
	}
}
