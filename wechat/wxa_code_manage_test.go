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
	_, err := client.WXA.ChangeVisitStatus(context.Background(), "token", "close")
	if err != nil {
		t.Errorf("WXA.ChangeVisitStatus retured err: %v", err)
	}
}

func TestWXAService_Commit(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &CommitRequest{
		TemplateID:      Int(0),
		ExtraJSON:       String("{\"extAppid\":\"\",\"ext\":{\"attr1\":\"value1\",\"attr2\":\"value2\"},\"extPages\":{\"index\":{},\"search/index\":{}},\"pages\":[\"index\",\"search/index\"],\"window\":{},\"networkTimeout\":{},\"tabBar\":{}}"),
		UserVersion:     String("V1.0"),
		UserDescription: String("test"),
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

func TestWXAService_GetAuditStatusByID(t *testing.T) {
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
	got, _, err := client.WXA.GetAuditStatusByID(context.Background(), "token", 1234567)
	if err != nil {
		t.Errorf("WXA.GetAuditStatusByID retured err: %v", err)
	}
	want := &AuditStatus{
		Status:     Int(1),
		Reason:     String("帐号信息不合规范"),
		Screenshot: String("xxx|yyy|zzz"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetAuditStatusByID got %+v, want %+v", got, want)
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
		Status:          Int(1),
		CreateTimestamp: Int(1517553721),
		GrayPercentage:  Int(8),
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
		Status:     Int(1),
		Reason:     String("帐号信息不合规范"),
		ScreenShot: String("xx|yy|zz"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetLatestAuditStatus got %+v, want %+v", got, want)
	}
}

func TestWXAService_GetPages(t *testing.T) {
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
	got, _, err := client.WXA.GetPages(context.Background(), "token")
	if err != nil {
		t.Errorf("WXA.GetPages retured err: %v", err)
	}
	want := &Pages{PageList: []*string{String("index"), String("page/list"), String("page/detail")}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WXA.GetPages got %+v, want %+v", got, want)
	}
}

func TestWXAService_GrayRelease(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &GrayReleaseRequest{GrayPercentage: Int(1)}
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
		Rest:         Int(0),
		Limit:        Int(0),
		SpeedupRest:  Int(0),
		SpeedupLimit: Int(0),
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
	_, err := client.WXA.SpeedupAudit(context.Background(), "token", 12345)
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
				Address:     String("index"),
				Tag:         String("学习 生活"),
				FirstClass:  String("文娱"),
				SecondClass: String("资讯"),
				FirstID:     Int(1),
				SecondID:    Int(2),
				Title:       String("首页"),
			},
			{
				Address:     String("page/logs/logs"),
				Tag:         String("学习 工作"),
				FirstClass:  String("教育"),
				SecondClass: String("学历教育"),
				ThirdClass:  String("高等"),
				FirstID:     Int(3),
				SecondID:    Int(4),
				ThirdID:     Int(5),
				Title:       String("日志"),
			},
		},
		PreviewInfo: &PreviewInfo{
			VideoIDs:   []*string{String("xxxx")},
			PictureIDs: []*string{String("xxxx"), String("yyyy"), String("zzzz")},
		},
		VersionDescription: String("blablabla"),
		FeedbackInfo:       String("blablabla"),
		FeedbackStuff:      String("xx|yy|zz"),
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
	want := &Audit{AuditID: Int(1234567)}
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
