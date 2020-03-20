package wechat

import (
	"context"
	"fmt"
)

// WXAService Wechat API docs: https://developers.weixin.qq.com/doc/
type WXAService service

// BindTesterRequest docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html
// #%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type BindTesterRequest struct {
	WechatID *string `json:"wechatid,omitempty"`
}

// BindTester docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type BindTester struct {
	UserString *string `json:"userstr,omitempty"`
}

// BindTester bind a tester for this app.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html
func (s *WXAService) BindTester(ctx context.Context, token string, r *BindTesterRequest) (*BindTester, *Response, error) {
	u := fmt.Sprintf("wxa/bind_tester?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}
	tester := new(BindTester)
	resp, err := s.client.Do(ctx, req, tester)
	if err != nil {
		return nil, resp, err
	}
	return tester, resp, nil
}
