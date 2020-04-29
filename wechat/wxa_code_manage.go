package wechat

import (
	"context"
	"fmt"
	"net/http"
)

// CommitRequest represents a request to commit code.
type CommitRequest struct {
	TemplateID  int    `json:"template_id"`
	ExtraJSON   string `json:"ext_json"`
	UserVersion string `json:"user_version"`
	UserDesc    string `json:"user_desc"`
}

// Commit commit code.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
func (s *WXAService) Commit(ctx context.Context, token string, r *CommitRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/commit?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Page represents a page response.
type Page struct {
	PageList []string `json:"page_list"`
}

// GetPage fetch app pages.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
func (s *WXAService) GetPage(ctx context.Context, token string) (*Page, *Response, error) {
	u := fmt.Sprintf("wxa/get_page?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	pages := new(Page)
	resp, err := s.client.Do(ctx, req, pages)
	if err != nil {
		return nil, resp, err
	}
	return pages, resp, nil
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

// Item represents a submit code item.
type Item struct {
	Address     string `json:"address,omitempty"`
	Tag         string `json:"tag,omitempty"`
	FirstClass  string `json:"first_class,omitempty"`
	SecondClass string `json:"second_class,omitempty"`
	ThirdClass  string `json:"third_class,omitempty"`
	FirstID     int    `json:"first_id,omitempty"`
	SecondID    int    `json:"second_id,omitempty"`
	ThirdID     int    `json:"third_id,omitempty"`
	Title       string `json:"title,omitempty"`
}

// PreviewInfo represents a submit audit preview info.
type PreviewInfo struct {
	VideoIDs   []string `json:"video_id_list,omitempty"`
	PictureIDs []string `json:"pic_id_list,omitempty"`
}

// SubmitAuditRequest represents a request to submit audit.
type SubmitAuditRequest struct {
	ItemList           []*Item      `json:"item_list,omitempty"`
	PreviewInfo        *PreviewInfo `json:"preview_info,omitempty"`
	VersionDescription string       `json:"version_desc,omitempty"`
	FeedbackInfo       string       `json:"feedback_info,omitempty"`
	FeedbackStuff      string       `json:"feedback_stuff,omitempty"`
}

// Audit represents a audit info.
type Audit struct {
	AuditID int `json:"auditid"`
}

// SubmitAudit submit cod to audit.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
func (s *WXAService) SubmitAudit(ctx context.Context, token string, r *SubmitAuditRequest) (*Audit, *Response, error) {
	u := fmt.Sprintf("wxa/submit_audit?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	audit := new(Audit)
	resp, err := s.client.Do(ctx, req, audit)
	if err != nil {
		return nil, resp, err
	}
	return audit, resp, nil
}

// AuditStatus represents a audit status.
type AuditStatus struct {
	Status     int    `json:"status"`
	Reason     string `json:"reason,omitempty"`
	Screenshot string `json:"screenshot,omitempty"`
	ScreenShot string `json:"ScreenShot,omitempty"` // Do not ask me why...
}

// GetAuditStatusRequest represents get audit status request.
type GetAuditStatusRequest struct {
	AuditID int `json:"auditid"`
}

// GetAuditStatus fetch audit status by auditID.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html
func (s *WXAService) GetAuditStatus(ctx context.Context, token string, r *GetAuditStatusRequest) (*AuditStatus, *Response, error) {
	u := fmt.Sprintf("wxa/get_auditstatus?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	status := new(AuditStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}
	return status, resp, nil
}

// GetLatestAuditStatus fetch latest audit status by auditID.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_latest_auditstatus.html
func (s *WXAService) GetLatestAuditStatus(ctx context.Context, token string) (*AuditStatus, *Response, error) {
	u := fmt.Sprintf("wxa/get_latest_auditstatus?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	status := new(AuditStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}
	return status, resp, nil
}

// UndoCodeAudit undo code audit.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/undocodeaudit.html
func (s *WXAService) UndoCodeAudit(ctx context.Context, token string) (*Response, error) {
	u := fmt.Sprintf("wxa/undocodeaudit?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Release release code.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/release.html
func (s *WXAService) Release(ctx context.Context, token string) (*Response, error) {
	u := fmt.Sprintf("wxa/release?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, struct{}{})
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// RevertCodeRelease revert code release.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertcoderelease.html
func (s *WXAService) RevertCodeRelease(ctx context.Context, token string) (*Response, error) {
	u := fmt.Sprintf("wxa/revertcoderelease?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// GrayReleaseRequest represents a request to grey release.
type GrayReleaseRequest struct {
	GrayPercentage int `json:"gray_percentage"`
}

// GrayRelease gray release.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/grayrelease.html
func (s *WXAService) GrayRelease(ctx context.Context, token string, r *GrayReleaseRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/grayrelease?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// GrayReleaseDetail represents a grey release detail.
type GrayReleaseDetail struct {
	GrayReleasePlan *GrayReleasePlan `json:"gray_release_plan,omitempty"`
}

// GrayReleasePlan represents grey release plan.
type GrayReleasePlan struct {
	Status          int `json:"status,omitempty"`
	CreateTimestamp int `json:"create_timestamp,omitempty"`
	GrayPercentage  int `json:"gray_percentage,omitempty"`
}

// GetGrayReleasePlan fetch gray release plan detail.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getgrayreleaseplan.html
func (s *WXAService) GetGrayReleasePlan(ctx context.Context, token string) (*GrayReleaseDetail, *Response, error) {
	u := fmt.Sprintf("wxa/getgrayreleaseplan?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	detail := new(GrayReleaseDetail)
	resp, err := s.client.Do(ctx, req, detail)
	if err != nil {
		return nil, resp, err
	}
	return detail, resp, nil
}

// RevertGrayRelease is to revert gray release.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertgrayrelease.html
func (s *WXAService) RevertGrayRelease(ctx context.Context, token string) (*Response, error) {
	u := fmt.Sprintf("wxa/revertgrayrelease?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// VisitStatus represents status for visit.
type VisitStatus struct {
	Action string `json:"action"`
}

// ChangeVisitStatusRequest represents change visit status request.
type ChangeVisitStatusRequest struct {
	Action string `json:"action"`
}

// ChangeVisitStatus change visit status.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/change_visitstatus.html
func (s *WXAService) ChangeVisitStatus(ctx context.Context, token string, r *ChangeVisitStatusRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/change_visitstatus?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Quota represents a quota count.
type Quota struct {
	Rest         int `json:"rest,omitempty"`
	Limit        int `json:"limit,omitempty"`
	SpeedupRest  int `json:"speedup_rest,omitempty"`
	SpeedupLimit int `json:"speedup_limit,omitempty"`
}

// QueryQuota fetch the quota info.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/query_quota.html
func (s *WXAService) QueryQuota(ctx context.Context, token string) (*Quota, *Response, error) {
	u := fmt.Sprintf("wxa/queryquota?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	quota := new(Quota)
	resp, err := s.client.Do(ctx, req, quota)
	if err != nil {
		return nil, resp, err
	}
	return quota, resp, nil
}

// SpeedupAuditRequest request for speedup audit
type SpeedupAuditRequest struct {
	AuditID int `json:"auditid"`
}

// SpeedupAudit to speedup audit.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/speedup_audit.html
func (s *WXAService) SpeedupAudit(ctx context.Context, token string, r *SpeedupAuditRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/speedupaudit?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodPost, u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
