package wechat

import (
	"context"
	"fmt"
)

// WXAService Wechat API docs: https://developers.weixin.qq.com/doc/
type WXAService service

// BindTesterRequest represents a request to bind a tester.
type BindTesterRequest struct {
	WechatID *string `json:"wechatid,omitempty"`
}

// BindTester represents a bind tester.
type BindTester struct {
	UserString *string `json:"userstr,omitempty"`
}

// BindTester bind a tester for this app.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Admin.html
func (s *WXAService) BindTester(ctx context.Context, token string, r *BindTesterRequest) (*BindTester, *Response, error) {
	u := fmt.Sprintf("wxa/bind_tester?access_token=%v", token)
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

// UnBindTesterRequest represents a request to unbind a tester.
type UnBindTesterRequest struct {
	UserString *string `json:"userstr,omitempty"`
	WechatID   *string `json:"wechatid,omitempty"`
}

// UnBindTester bind a tester for this app.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/unbind_tester.html
func (s *WXAService) UnBindTester(ctx context.Context, token string, r *UnBindTesterRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/unbind_tester?access_token=%v", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Testers represents a request a bind tester list.
type Testers struct {
	Members []*BindTester `json:"members,omitempty"`
}

// GetTesters get tester list.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/memberauth.html
func (s *WXAService) GetTesters(ctx context.Context, token string) (*Testers, *Response, error) {
	u := fmt.Sprintf("wxa/memberauth?access_token=%v", token)
	payload := map[string]string{
		"action": "get_experiencer",
	}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}
	testers := new(Testers)
	resp, err := s.client.Do(ctx, req, testers)
	if err != nil {
		return nil, resp, err
	}
	return testers, resp, nil
}

// ShowWXAItem represents a item of show wxa.
type ShowWXAItem struct {
	CanOpen   *int    `json:"can_open,omitempty"`
	IsOpen    *int    `json:"is_open,omitempty"`
	AppID     *string `json:"app_id,omitempty"`
	Nickname  *string `json:"nickname,omitempty"`
	HeadImage *string `json:"headimg,omitempty"`
}

// GetShowWXAItem get show wxa item
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getshowwxaitem.html
func (s *WXAService) GetShowWXAItem(ctx context.Context, token string) (*ShowWXAItem, *Response, error) {
	u := fmt.Sprintf("wxa/getshowwxaitem?access_token=%v", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	showWXAItem := new(ShowWXAItem)
	resp, err := s.client.Do(ctx, req, showWXAItem)
	if err != nil {
		return nil, resp, err
	}
	return showWXAItem, resp, nil
}

// WXAMPLink link.
type WXAMPLink struct {
	Nickname  *string `json:"nickname,omitempty"`
	AppID     *string `json:"appid,omitempty"`
	HeadImage *string `json:"headimg,omitempty"`
}

// WXAMPLinks link list.
type WXAMPLinks struct {
	BIZInfoList []*WXAMPLink `json:"biz_info_list,omitempty"`
	TotalNum    int          `json:"total_num,omitempty"`
}

// GetShowWXAItem get show wxa item
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/getwxamplinkforshow.html
func (s *WXAService) GetWXAMpLinkForShow(ctx context.Context, token string, page, num int) (*WXAMPLinks, *Response, error) {
	u := fmt.Sprintf("wxa/getwxamplinkforshow?access_token=%v&page=%d&num=%d", token, page, num)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	wxaMpLinks := new(WXAMPLinks)
	resp, err := s.client.Do(ctx, req, wxaMpLinks)
	if err != nil {
		return nil, resp, err
	}
	return wxaMpLinks, resp, nil
}

// UpdateShowWXAItemRequest update show wxa.
type UpdateShowWXAItemRequest struct {
	WXASubscribeBIZFlag int    `json:"wxa_subscribe_biz_flag,omitempty"`
	AppID               string `json:"appid,omitempty"`
}

// UpdateShowWXAItem get show wxa item
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/subscribe_component/updateshowwxaitem.html
func (s *WXAService) UpdateShowWXAItem(ctx context.Context, token string, r *UpdateShowWXAItemRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/updateshowwxaitem?access_token=%v", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
