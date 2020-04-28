package wechat

import (
	"context"
	"fmt"
	"net/http"
)

type ModifyDomainRequest struct {
	Action          string   `json:"action"`
	RequestDomain   []string `json:"requestdomain,omitempty"`
	WSRequestDomain []string `json:"wsrequestdomain,omitempty"`
	UploadDomain    []string `json:"uploaddomain,omitempty"`
	DownloadDomain  []string `json:"downloaddomain,omitempty"`
}

type Domain struct {
	RequestDomain   []string `json:"requestdomain,omitempty"`
	WSRequestDomain []string `json:"wsrequestdomain,omitempty"`
	UploadDomain    []string `json:"uploaddomain,omitempty"`
	DownloadDomain  []string `json:"downloaddomain,omitempty"`
}

// ModifyDomain server address configuration.
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html
func (s *WXAService) ModifyDomain(ctx context.Context, token string, r *ModifyDomainRequest) (*Domain, *Response, error) {
	u := fmt.Sprintf("wxa/modify_domain?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, nil, err
	}
	domain := new(Domain)
	resp, err := s.client.Do(ctx, req, domain)
	if err != nil {
		return nil, resp, err
	}
	return domain, resp, nil
}

type SetWebViewDomainRequest struct {
	Action        string   `json:"action,omitempty"`
	WebViewDomain []string `json:"webviewdomain,omitempty"`
}

// ModifyDomain server address configuration.
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Server_Address_Configuration.html
func (s *WXAService) SetWebViewDomain(ctx context.Context, token string, r *SetWebViewDomainRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/setwebviewdomain?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
