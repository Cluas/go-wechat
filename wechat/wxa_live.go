package wechat

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// GetLiveInfoRequest represents request of get live info.
type GetLiveInfoRequest struct {
	Action string `json:"action,omitempty"`
	RoomID int    `json:"room_id,omitempty"`
	Start  int    `json:"start"`
	Limit  int    `json:"limit"`
}

// LiveReplay struct
type LiveReplay struct {
	ExpireTime time.Time `json:"expire_time"`
	CreateTime time.Time `json:"create_time"`
	MediaURL   string    `json:"media_url"`
}

// Good represents business good
type Good struct {
	CoverImg  string `json:"cover_img"`
	URL       string `json:"url"`
	Price     int    `json:"price"`
	Price2    int    `json:"price2"`
	PriceType int    `json:"price_type"`
	Name      string `json:"name"`
}

// RoomInfo represents live room info
type RoomInfo struct {
	Name       string  `json:"name"`
	RoomID     int     `json:"roomid"`
	CoverImg   string  `json:"cover_img"`
	LiveStatus int     `json:"live_status"`
	StartTime  int64   `json:"start_time"`
	EndTime    int64   `json:"end_time"`
	AnchorName string  `json:"anchor_name"`
	ShareImg   string  `json:"share_img"` // response diff with doc
	Goods      []*Good `json:"goods"`
}

// LiveInfo represents get live info response
type LiveInfo struct {
	RoomInfo   []*RoomInfo   `json:"room_info,omitempty"`
	LiveReplay []*LiveReplay `json:"live_replay,omitempty"`
	Total      int           `json:"total"`
}

// GetLiveInfo represents get live info.
// Wechat API docs:
// https://developers.weixin.qq.com/miniprogram/dev/framework/liveplayer/live-player-plugin.html
func (s *WXAService) GetLiveInfo(ctx context.Context, token string, r *GetLiveInfoRequest) (*LiveInfo, *Response, error) {
	u := fmt.Sprintf("wxa/business/getliveinfo?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, nil, err
	}
	liveInfo := new(LiveInfo)
	resp, err := s.client.Do(ctx, req, liveInfo)
	if err != nil {
		return nil, resp, err
	}
	return liveInfo, resp, nil
}
