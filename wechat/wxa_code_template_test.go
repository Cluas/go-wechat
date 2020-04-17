package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_AddDraftToTemplate(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/addtotemplate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							}`)
	})
	_, err := client.WXA.AddDraftToTemplate(context.Background(), "token", 0)
	if err != nil {
		t.Errorf("WXA.AddDraftToTemplate retured err: %v", err)
	}
}
func TestWXAService_DeleteTemplateByID(t *testing.T) {
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
	_, err := client.WXA.DeleteTemplateByID(context.Background(), "token", 0)
	if err != nil {
		t.Errorf("WXA.DeleteTemplateByID retured err: %v", err)
	}

}

func TestWXAService_GetTemplateDrafts(t *testing.T) {
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
	got, _, err := client.WXA.GetTemplateDrafts(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetTemplateDrafts retured err: %v", err)
	}
	want := &TemplateDrafts{DraftList: []*Draft{
		{
			CreateTime:      Int(1488965944),
			UserVersion:     String("VVV"),
			UserDescription: String("AAS"),
			DraftID:         Int(0),
		},
		{
			CreateTime:      Int(1504790906),
			UserVersion:     String("11"),
			UserDescription: String("111111"),
			DraftID:         Int(4),
		},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetTemplateDrafts got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetTemplates(t *testing.T) {
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
	got, _, err := client.WXA.GetTemplates(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetTemplates retured err: %v", err)
	}
	want := &Templates{TemplateList: []*Template{
		{
			CreateTime:      Int(1488965944),
			UserVersion:     String("VVV"),
			UserDescription: String("AAS"),
			TemplateID:      Int(0),
		},
		{
			CreateTime:      Int(1504790906),
			UserVersion:     String("11"),
			UserDescription: String("111111"),
			TemplateID:      Int(4),
		},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetTemplates got %+v, want %+v", got, want)
	}
}
