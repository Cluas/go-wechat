package wechat

import (
	"context"
	"fmt"
	"net/http"
)

// ComponentService Wechat API docs: https://developers.weixin.qq.com/doc/
type ComponentService service

// APIComponentTokenRequest represents a request to get a token.
type APIComponentTokenRequest struct {
	ComponentAppID        string `json:"component_app_id"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

// Token represents a API token on a component.
type Token struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

// APIComponentToken fetch a new api component token.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_access_token.html
func (s *ComponentService) APIComponentToken(ctx context.Context, r *APIComponentTokenRequest) (*Token, *Response, error) {
	u := "cgi-bin/component/api_component_token"
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, nil, err
	}
	token := new(Token)
	resp, err := s.client.Do(ctx, req, token)
	if err != nil {
		return nil, resp, err
	}
	return token, resp, nil
}

// FastRegisterWeAppRequest represents a request to create a mini program.
type FastRegisterWeAppRequest struct {
	Name               string `json:"name"`
	Code               string `json:"code"`
	CodeType           int    `json:"code_type"`
	LegalPersonaWechat string `json:"legal_persona_wechat"`
	LegalPersonaName   string `json:"legal_persona_name"`
	ComponentPhone     string `json:"component_phone"`
}

// FastRegisterWeApp create a new mini program.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Fast_Registration_Interface_document.html
func (s *ComponentService) FastRegisterWeApp(ctx context.Context, token string, r *FastRegisterWeAppRequest) (*Response, error) {
	u := fmt.Sprintf("cgi-bin/component/fastregisterweapp?action=create&component_access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
