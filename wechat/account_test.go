package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBasicInfo_marshal(t *testing.T) {
	testJSONMarshal(t, &BasicInfo{}, "{}")

	b := &BasicInfo{
		AppID:          String("app_id"),
		AccountType:    Int(1),
		PrincipalType:  Int(1),
		PrincipalName:  String("principal_name"),
		RealNameStatus: String("verify"),
		WXVerifyInfo: &WXVerifyInfo{
			QualificationVerify: Bool(true),
			NamingVerify:        Bool(true),
			AnnualReview:        Bool(true),
		},
		SignatureInfo: &SignatureInfo{
			Signature:       String("signature"),
			ModifyUsedCount: Int(1),
			ModifyQuota:     Int(1),
		},
		HeadImageInfo: &HeadImageInfo{
			HeadImageURL:    String("url"),
			ModifyUsedCount: Int(1),
			ModifyQuota:     Int(1),
		},
	}
	want := `
		{
			"app_id": "app_id",
			"account_type": 1,
			"principal_type": 1,
			"principal_name": "principal_name",
			"realname_status": "verify",
			"wx_verify_info": {
				"qualification_verify": true,
				"naming_verify": true,
				"annual_review": true
			},
			"signature_info": {
				"signature": "signature",
				"modify_used_count": 1,
				"modify_quota": 1
			},
			"head_image_info": {
				"head_image_url": "url",
				"modify_used_count": 1,
				"modify_quota": 1
			}
		}
	`
	testJSONMarshal(t, b, want)
}

func TestAccountService_GetBasicInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/cgi-bin/account/getaccountbasicinfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"app_id":"app_id", "account_type":1, "principal_type":1, "principal_name": "principal_name", "realname_status":"verify"}`)
	})
	info, _, err := client.Account.GetBasicInfo(context.Background(), "token")
	if err != nil {
		t.Errorf("Account.GetBasicInfo returned error: %v", err)
	}
	want := &BasicInfo{
		AppID:          String("app_id"),
		AccountType:    Int(1),
		PrincipalType:  Int(1),
		PrincipalName:  String("principal_name"),
		RealNameStatus: String("verify"),
	}
	if !reflect.DeepEqual(info, want) {
		t.Errorf("Account.GetBasicInfo returned %+v, want %+v", info, want)
	}
}
