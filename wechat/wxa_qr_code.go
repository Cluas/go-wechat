package wechat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateQRCodeRequest represents request of create qr code.
type CreateQRCodeRequest struct {
	Path  string `json:"path"`
	Width int    `json:"width,omitempty"`
}

// CreateWXAQRCode
// Wechat API docs:
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (s *WXAService) CreateWXAQRCode(ctx context.Context, token string, r *CreateQRCodeRequest) (*Response, error) {
	u := fmt.Sprintf("cgi-bin/wxaapp/createwxaqrcode?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	s.client.clientMu.Lock()
	defer s.client.clientMu.Unlock()
	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// LineColor line color
type LineColor struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// GetWXACodeRequest represents request of get qr code.
type GetWXACodeRequest struct {
	Path      string     `json:"path"`
	Width     int        `json:"width,omitempty"`
	AutoColor bool       `json:"auto_color,omitempty"`
	LineColor *LineColor `json:"line_color,omitempty"`
	IsHyaline bool       `json:"is_hyaline,omitempty"`
}

// GetWXACode
// Wechat API docs:
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func (s *WXAService) GetWXACode(ctx context.Context, token string, r *GetWXACodeRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/getwxacode?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	s.client.clientMu.Lock()
	defer s.client.clientMu.Unlock()
	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// GetWXACodeUnlimitRequest represents request of get qr code.
type GetWXACodeUnlimitRequest struct {
	Scene     string     `json:"scene"`
	Path      string     `json:"path,omitempty"`
	Width     int        `json:"width,omitempty"`
	AutoColor bool       `json:"auto_color,omitempty"`
	LineColor *LineColor `json:"line_color,omitempty"`
	IsHyaline bool       `json:"is_hyaline,omitempty"`
}

// GetWXACodeUnlimit
// Wechat API docs:
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (s *WXAService) GetWXACodeUnlimit(ctx context.Context, token string, r *GetWXACodeUnlimitRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/getwxacodeunlimit?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	s.client.clientMu.Lock()
	defer s.client.clientMu.Unlock()
	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// GetQrCode fetch qr_code.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_qrcode.html
func (s *WXAService) GetQrCode(ctx context.Context, token, path string) (*Response, error) {
	u := fmt.Sprintf("wxa/get_qrcode?access_token=%v&path=%s", token, path)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	s.client.clientMu.Lock()
	defer s.client.clientMu.Unlock()
	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}
