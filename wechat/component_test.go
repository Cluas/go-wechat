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

	req := &TokenRequest{
		ComponentAppID:        String("component_app_id"),
		ComponentAppSecret:    String("component_app_secret"),
		ComponentVerifyTicket: String("component_verify_ticket"),
	}

	mux.HandleFunc("/cgi-bin/component/api_component_token", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "component_access_token": "61W3mEpU66027wgNZ_MhGHNQDHnFATkDa9-2llqrMBjUwxRSNPbVsMmyD-yq8wZETSoE5NQgecigDrSHkPtIYA",
							  "expires_in": 7200
							}`)
	})
	got, _, err := client.Component.GetToken(context.Background(), req)
	if err != nil {
		t.Errorf("Component.GetToken returned error: %v", err)
	}
	want := &Token{
		ComponentAccessToken: String("61W3mEpU66027wgNZ_MhGHNQDHnFATkDa9-2llqrMBjUwxRSNPbVsMmyD-yq8wZETSoE5NQgecigDrSHkPtIYA"),
		ExpiresIn:            Int(7200),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Component.GetToken returned %+v, want %+v", got, want)
	}
}

func TestComponentService_CreateMiniProgram(t *testing.T) {
	client, mux, _, tearDown := setup()
	defer tearDown()

	req := &CreateMiniProgramRequest{
		Name:               String("tencent"),
		Code:               String("123"),
		CodeType:           Int(1),
		LegalPersonaWechat: String("123"),
		LegalPersonaName:   String("pony"),
		ComponentPhone:     String("1234567"),
	}
	mux.HandleFunc("/cgi-bin/component/fastregisterweapp?action=create", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
							  "errcode": 0, 
							  "errmsg": "OK"
							}`)
	})
	_, err := client.Component.CreateMiniProgram(context.Background(), "token", req)
	if err != nil {
		t.Errorf("Component.CreateMiniProgram returend error: %v", err)
	}
}
