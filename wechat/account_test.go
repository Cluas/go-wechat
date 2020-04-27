package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBasicInfo_marshal(t *testing.T) {
	testJSONMarshal(t, &AccountBasicInfo{}, "{}")

	b := &AccountBasicInfo{
		AppID:          "app_id",
		AccountType:    1,
		PrincipalType:  1,
		PrincipalName:  "principal_name",
		RealNameStatus: "verify",
		WXVerifyInfo: &WXVerifyInfo{
			QualificationVerify: true,
			NamingVerify:        true,
			AnnualReview:        true,
		},
		SignatureInfo: &SignatureInfo{
			Signature:       "signature",
			ModifyUsedCount: 1,
			ModifyQuota:     1,
		},
		HeadImageInfo: &HeadImageInfo{
			HeadImageURL:    "url",
			ModifyUsedCount: 1,
			ModifyQuota:     1,
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

func TestAccountService_GetAccountBasicInfo(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	mux.HandleFunc("/cgi-bin/account/getaccountbasicinfo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"app_id":"app_id", "account_type":1, "principal_type":1, "principal_name": "principal_name", "realname_status":"verify"}`)
	})
	info, _, err := client.Account.GetAccountBasicInfo(context.Background(), "token")
	if err != nil {
		t.Errorf("Account.GetBasicInfo returned error: %v", err)
	}
	want := &AccountBasicInfo{
		AppID:          "app_id",
		AccountType:    1,
		PrincipalType:  1,
		PrincipalName:  "principal_name",
		RealNameStatus: "verify",
	}
	if !reflect.DeepEqual(info, want) {
		t.Errorf("Account.GetBasicInfo returned %+v, want %+v", info, want)
	}
}
