package wechat

import (
	"context"
	"fmt"
	"net/http"
)

// PluginRequest represents plugin request
type PluginRequest struct {
	Action      string `json:"action"`
	PluginAppID string `json:"plugin_appid,omitempty"`
	UserVersion string `json:"user_version,omitempty"`
}

// Plugin represents plugin struct
type Plugin struct {
	AppID      string `json:"appid"`
	Status     int    `json:"status"`
	Nickname   string `json:"nickname"`
	HeadImgURL string `json:"headimgurl"`
}

// PluginResponse represents plugin response
type PluginResponse struct {
	PluginList []*Plugin `json:"plugin_list"`
	ApplyList  []*Plugin `json:"apply_list"`
}

// Plugin represents manage wxa plugins.
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Plug-ins_Management.html
func (s *WXAService) Plugin(ctx context.Context, token string, r *PluginRequest) (*PluginResponse, *Response, error) {
	u := fmt.Sprintf("wxa/plugin?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, nil, err
	}
	pluginResp := new(PluginResponse)
	resp, err := s.client.Do(ctx, req, pluginResp)
	if err != nil {
		return nil, resp, err
	}
	return pluginResp, resp, nil
}
