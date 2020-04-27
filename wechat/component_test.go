package wechat

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestComponentService_GetToken(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &APIComponentTokenRequest{
		ComponentAppID:        "component_app_id",
		ComponentAppSecret:    "component_app_secret",
		ComponentVerifyTicket: "component_verify_ticket",
	}

	mux.HandleFunc("/cgi-bin/component/api_component_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "component_access_token": "61W3mEpU66027wgNZ_MhGHNQDHnFATkDa9-2llqrMBjUwxRSNPbVsMmyD-yq8wZETSoE5NQgecigDrSHkPtIYA",
							  "expires_in": 7200
							}`)
	})
	got, _, err := client.Component.APIComponentToken(context.Background(), req)
	if err != nil {
		t.Errorf("Component.APIComponentToken returned error: %v", err)
	}
	want := &Token{
		ComponentAccessToken: "61W3mEpU66027wgNZ_MhGHNQDHnFATkDa9-2llqrMBjUwxRSNPbVsMmyD-yq8wZETSoE5NQgecigDrSHkPtIYA",
		ExpiresIn:            7200,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Component.APIComponentToken returned %+v, want %+v", got, want)
	}
}

func TestComponentService_CreateMiniProgram(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &FastRegisterWeAppRequest{
		Name:               "tencent",
		Code:               "123",
		CodeType:           1,
		LegalPersonaWechat: "123",
		LegalPersonaName:   "pony",
		ComponentPhone:     "1234567",
	}
	mux.HandleFunc("/cgi-bin/component/fastregisterweapp?action=create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0, 
							  "errmsg": "OK"
							}`)
	})
	_, err := client.Component.FastRegisterWeApp(context.Background(), "token", req)
	if err != nil {
		t.Errorf("Component.FastRegisterWeApp returend error: %v", err)
	}
}
