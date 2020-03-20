package wechat

import (
	"context"
	"fmt"
)

// Wechat API docs: https://developers.weixin.qq.com/doc/
type ComponentService service

// TokenRequest represents a request to get a token.
type TokenRequest struct {
	ComponentAppID        *string `json:"component_app_id"`
	ComponentAppSecret    *string `json:"component_appsecret"`
	ComponentVerifyTicket *string `json:"component_verify_ticket"`
}

// Token represents a API token on a component.
type Token struct {
	ComponentAccessToken *string `json:"component_access_token"`
	ExpireIn             *int    `json:"expire_in"`
}

// GetToken fetch a new api component token.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/api/component_access_token.html
func (s *ComponentService) GetToken(ctx context.Context, r *TokenRequest) (*Token, *Response, error) {
	u := "cgi-bin/component/api_component_token"
	req, err := s.client.NewRequest("POST", u, r)
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

// CreateMiniProgramRequest represents a request to create a mini program
type CreateMiniProgramRequest struct {
	Name               *string `json:"name"`
	Code               *string `json:"code"`
	CodeType           *int    `json:"code_type"`
	LegalPersonaWechat *string `json:"legal_persona_wechat"`
	LegalPersonaName   *string `json:"legal_persona_name"`
	ComponentPhone     *string `json:"component_phone"`
}

// CreateMiniProgram create a new mini program.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Fast_Registration_Interface_document.html
func (s *ComponentService) CreateMiniProgram(ctx context.Context, token string, r *TokenRequest) (*Response, error) {
	u := fmt.Sprintf("cgi-bin/component/fastregisterweapp?action=create&component_access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
