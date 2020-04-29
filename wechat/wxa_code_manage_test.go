package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWXAService_ChangeVisitStatus(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/change_visitstatus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.ChangeVisitStatus(context.Background(), "token", &ChangeVisitStatusRequest{Action: "close"})
	if err != nil {
		t.Errorf("WXA.ChangeVisitStatus retured err: %v", err)
	}
}

func TestWXAService_Commit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &CommitRequest{
		TemplateID:  0,
		ExtraJSON:   "{\"extAppid\":\"\",\"ext\":{\"attr1\":\"value1\",\"attr2\":\"value2\"},\"extPages\":{\"index\":{},\"search/index\":{}},\"pages\":[\"index\",\"search/index\"],\"window\":{},\"networkTimeout\":{},\"tabBar\":{}}",
		UserVersion: "V1.0",
		UserDesc:    "test",
	}

	mux.HandleFunc("/wxa/commit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})

	_, err := client.WXA.Commit(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.Commit retured err: %v", err)
	}
}

func TestWXAService_GetAuditStatus(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/get_auditstatus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "status": 1,
							  "reason": "帐号信息不合规范",
							  "screenshot": "xxx|yyy|zzz"
							}`)
	})
	got, _, err := client.WXA.GetAuditStatus(context.Background(), "token", &GetAuditStatusRequest{AuditID: 0})
	if err != nil {
		t.Errorf("WXA.GetAuditStatus retured err: %v", err)
	}
	want := &AuditStatus{
		Status:     1,
		Reason:     "帐号信息不合规范",
		Screenshot: "xxx|yyy|zzz",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetAuditStatus got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetGrayReleasePlan(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/getgrayreleaseplan", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "gray_release_plan": {
								"status": 1,
								"create_timestamp": 1517553721,
								"gray_percentage": 8
							  }
							}`)
	})
	got, _, err := client.WXA.GetGrayReleasePlan(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetGrayReleasePlan retured err: %v", err)
	}
	want := &GrayReleaseDetail{GrayReleasePlan: &GrayReleasePlan{
		Status:          1,
		CreateTimestamp: 1517553721,
		GrayPercentage:  8,
	}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetGrayReleasePlan got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetLatestAuditStatus(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/get_latest_auditstatus", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "auditid": "1234567",
							  "status": 1,
							  "reason": "帐号信息不合规范",
							  "ScreenShot": "xx|yy|zz"
							}`)
	})
	got, _, err := client.WXA.GetLatestAuditStatus(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetLatestAuditStatus retured err: %v", err)
	}
	want := &AuditStatus{
		Status:     1,
		Reason:     "帐号信息不合规范",
		ScreenShot: "xx|yy|zz",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetLatestAuditStatus got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetPage(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/get_page", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "page_list": ["index", "page/list", "page/detail"]
							}`)
	})
	got, _, err := client.WXA.GetPage(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetPage retured err: %v", err)
	}
	want := &Page{PageList: []string{"index", "page/list", "page/detail"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetPage got %+v, want %+v", got, want)
	}
}

func TestWXAService_GrayRelease(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &GrayReleaseRequest{GrayPercentage: 1}
	mux.HandleFunc("/wxa/grayrelease", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.GrayRelease(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.GrayRelease retured err: %v", err)
	}
}

func TestWXAService_QueryQuota(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/queryquota", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
							  "rest": 0,
							  "limit": 0,
							  "speedup_rest": 0,
							  "speedup_limit": 0
							}`)
	})
	got, _, err := client.WXA.QueryQuota(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.QueryQuota retured err: %v", err)
	}
	want := &Quota{
		Rest:         0,
		Limit:        0,
		SpeedupRest:  0,
		SpeedupLimit: 0,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.QueryQuota got %+v, want %+v", got, want)
	}
}

func TestWXAService_Release(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/release", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.Release(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.Release retured err: %v", err)
	}
}

func TestWXAService_RevertCodeRelease(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/revertcoderelease", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.RevertCodeRelease(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.RevertCodeRelease retured err: %v", err)
	}
}

func TestWXAService_RevertGrayRelease(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/revertgrayrelease", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.RevertGrayRelease(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.RevertGrayRelease retured err: %v", err)
	}
}

func TestWXAService_SpeedupAudit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/speedupaudit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})
	_, err := client.WXA.SpeedupAudit(context.Background(), "token", &SpeedupAuditRequest{AuditID: 12345})
	if err != nil {
		t.Errorf("WXA.SpeedupAudit retured err: %v", err)
	}
}

func TestWXAService_SubmitAudit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &SubmitAuditRequest{
		ItemList: []*Item{
			{
				Address:     "index",
				Tag:         "学习 生活",
				FirstClass:  "文娱",
				SecondClass: "资讯",
				FirstID:     1,
				SecondID:    2,
				Title:       "首页",
			},
			{
				Address:     "page/logs/logs",
				Tag:         "学习 工作",
				FirstClass:  "教育",
				SecondClass: "学历教育",
				ThirdClass:  "高等",
				FirstID:     3,
				SecondID:    4,
				ThirdID:     5,
				Title:       "日志",
			},
		},
		PreviewInfo: &PreviewInfo{
			VideoIDs:   []string{"xxxx"},
			PictureIDs: []string{"xxxx", "yyyy", "zzzz"},
		},
		VersionDescription: "blablabla",
		FeedbackInfo:       "blablabla",
		FeedbackStuff:      "xx|yy|zz",
	}
	mux.HandleFunc("/wxa/submit_audit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok",
                              "auditid": 1234567
							}`)
	})
	got, _, err := client.WXA.SubmitAudit(context.Background(), "token", req)
	if err != nil {
		t.Errorf("WXA.SubmitAudit retured err: %v", err)
	}
	want := &Audit{AuditID: 1234567}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.SubmitAudit got %+v, want %+v", got, want)
	}

}

func TestWXAService_UndoCodeAudit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/wxa/undocodeaudit", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
							  "errcode": 0,
							  "errmsg": "ok"
							}`)
	})

	_, err := client.WXA.UndoCodeAudit(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.UndoCodeAudit retured err: %v", err)
	}
}
