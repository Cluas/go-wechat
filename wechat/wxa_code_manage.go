package wechat

import (
	"context"
	"fmt"
	"time"
)

// CommitRequest docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
// #%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type CommitRequest struct {
	TemplateID      *int    `json:"template_id,omitempty"`
	ExtraJSON       *string `json:"ext_json,omitempty"`
	UserVersion     *string `json:"user_version,omitempty"`
	UserDescription *string `json:"user_desc,omitempty"`
}

// Commit commit code.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
func (s *WXAService) Commit(ctx context.Context, token string, r *CommitRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/commit?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Pages docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type Pages struct {
	PageList []*string `json:"page_list,omitempty"`
}

// GetPages fetch app pages.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
func (s *WXAService) GetPages(ctx context.Context, token string) (*Pages, *Response, error) {
	u := fmt.Sprintf("wxa/get_page?access_token=%s", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	pages := new(Pages)
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
	u := fmt.Sprintf("wxa/get_qrcode?access_token=%s&path=%s", token, path)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	s.client.clientMu.Lock()
	defer s.client.clientMu.Unlock()
	resp, err := s.client.Do(ctx, req, nil)
	return resp, err
}

// Item docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
// #%E5%AE%A1%E6%A0%B8%E9%A1%B9%E8%AF%B4%E6%98%8E
type Item struct {
	Address     *string `json:"address,omitempty"`
	Tag         *string `json:"tag,omitempty"`
	FirstClass  *string `json:"first_class,omitempty"`
	SecondClass *string `json:"second_class,omitempty"`
	ThirdClass  *string `json:"third_class,omitempty"`
	FirstID     *int    `json:"first_id,omitempty"`
	SecondID    *int    `json:"second_id,omitempty"`
	ThirdID     *int    `json:"third_id,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// PreviewInfo docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
// #%E9%A2%84%E8%A7%88%E4%BF%A1%E6%81%AF%E8%AF%B4%E6%98%8E
type PreviewInfo struct {
	VideoIDs   []*int `json:"video_id_list,omitempty"`
	PictureIDs []*int `json:"pic_id_list,omitempty"`
}

// SubmitAuditRequest docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
// #%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type SubmitAuditRequest struct {
	ItemList           []*Item      `json:"item_list,omitempty"`
	PreviewInfo        *PreviewInfo `json:"preview_info,omitempty"`
	VersionDescription *string      `json:"version_desc,omitempty"`
	FeedbackInfo       *string      `json:"feedback_info,omitempty"`
	FeedbackStuff      *string      `json:"feedback_stuff,omitempty"`
}

// Audit struct
type Audit struct {
	AuditID *int `json:"auditid,omitempty"`
}

// SubmitAudit submit cod to audit.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
func (s *WXAService) SubmitAudit(ctx context.Context, token string, r *SubmitAuditRequest) (*Audit, *Response, error) {
	u := fmt.Sprintf("wxa/submit_audit?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, r)
	audit := new(Audit)
	resp, err := s.client.Do(ctx, req, audit)
	if err != nil {
		return nil, resp, err
	}
	return audit, resp, nil
}

// AuditStatus docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type AuditStatus struct {
	Status     *int    `json:"status,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	Screenshot *string `json:"screenshot,omitempty"`
	ScreenShot *string `json:"ScreenShot,omitempty"` // Do not ask me why...
}

// GetAuditStatusByID fetch audit status by auditID.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_auditstatus.html
func (s *WXAService) GetAuditStatusByID(ctx context.Context, token string, auditID int) (*AuditStatus, *Response, error) {
	u := fmt.Sprintf("wxa/get_auditstatus?access_token=%s", token)
	payload := &Audit{Int(auditID)}
	req, err := s.client.NewRequest("POST", u, payload)
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
	u := fmt.Sprintf("wxa/get_latest_auditstatus?access_token=%s", token)
	req, err := s.client.NewRequest("GET", u, nil)
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
	u := fmt.Sprintf("wxa/undocodeaudit?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, nil)
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
	u := fmt.Sprintf("wxa/release?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, nil)
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
	u := fmt.Sprintf("wxa/revertcoderelease?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// GrayReleaseRequest docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/grayrelease.html
// #%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type GrayReleaseRequest struct {
	GrayPercentage *int `json:"gray_percentage,omitempty"`
}

// GrayRelease gray release.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/grayrelease.html
func (s *WXAService) GrayRelease(ctx context.Context, token string, r *GrayReleaseRequest) (*Response, error) {
	u := fmt.Sprintf("wxa/grayrelease?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// GrayReleaseDetail docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getgrayreleaseplan.html
// #%E5%88%86%E9%98%B6%E6%AE%B5%E5%8F%91%E5%B8%83%E8%AE%A1%E5%88%92%E8%AF%A6%E6%83%85
type GrayReleaseDetail struct {
	GrayReleasePlan *GrayReleasePlan `json:"gray_release_plan,omitempty"`
}

// GrayReleasePlan struct
type GrayReleasePlan struct {
	Status          *int       `json:"status,omitempty"`
	CreateTimestamp *time.Time `json:"create_timestamp,omitempty"`
	GrayPercentage  *int       `json:"gray_percentage,omitempty"`
}

// GetGrayReleasePlan fetch gray release plan detail.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/getgrayreleaseplan.html
func (s *WXAService) GetGrayReleasePlan(ctx context.Context, token string) (*GrayReleaseDetail, *Response, error) {
	u := fmt.Sprintf("wxa/getgrayreleaseplan?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, nil)
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
	u := fmt.Sprintf("wxa/revertgrayrelease?access_token=%s", token)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// VisitStatus docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/change_visitstatus.html
type VisitStatus struct {
	Action *string `json:"action,omitempty"`
}

// ChangeVisitStatus change visit status.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/change_visitstatus.html
func (s *service) ChangeVisitStatus(ctx context.Context, token string, action string) (*Response, error) {
	u := fmt.Sprintf("wxa/change_visitstatus?access_token=%s", token)
	payload := &VisitStatus{Action: String(action)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Quota docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/query_quota.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type Quota struct {
	Rest         *int `json:"rest,omitempty"`
	Limit        *int `json:"limit,omitempty"`
	SpeedupRest  *int `json:"speedup_rest,omitempty"`
	SpeedupLimit *int `json:"speedup_limit,omitempty"`
}

// QueryQuota fetch the quota info.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/query_quota.html
func (s *WXAService) QueryQuota(ctx context.Context, token string) (*Response, error) {
	u := fmt.Sprintf("wxa/queryquota?access_token=%s", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// SpeedupAudit to speedup audit.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/speedup_audit.html
func (s *WXAService) SpeedupAudit(ctx context.Context, token string, auditID int) (*Response, error) {
	u := fmt.Sprintf("wxa/speedupaudit?access_token=%s", token)
	payload := &Audit{Int(auditID)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
