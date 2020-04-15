package wechat

import (
	"context"
	"fmt"
	"time"
)

// AccountService Wechat API docs: https://developers.weixin.qq.com/doc/
type AccountService service

// WXVerifyInfo represents wechat verify info.
type WXVerifyInfo struct {
	QualificationVerify   *bool      `json:"qualification_verify,omitempty"`
	NamingVerify          *bool      `json:"naming_verify,omitempty"`
	AnnualReview          *bool      `json:"annual_review,omitempty"`
	AnnualReviewBeginTime *time.Time `json:"annual_review_begin_time,omitempty"`
	AnnualReviewEndTime   *time.Time `json:"annual_review_end_time,omitempty"`
}

// SignatureInfo represents wechat signature info.
type SignatureInfo struct {
	Signature       *string `json:"signature,omitempty"`
	ModifyUsedCount *int    `json:"modify_used_count,omitempty"`
	ModifyQuota     *int    `json:"modify_quota,omitempty"`
}

// HeadImageInfo represents wechat head image info.
type HeadImageInfo struct {
	HeadImageURL    *string `json:"head_image_url,omitempty"`
	ModifyUsedCount *int    `json:"modify_used_count,omitempty"`
	ModifyQuota     *int    `json:"modify_quota,omitempty"`
}

// BasicInfo represents wechat basic info.
type BasicInfo struct {
	AppID          *string        `json:"app_id,omitempty"`
	AccountType    *int           `json:"account_type,omitempty"`
	PrincipalType  *int           `json:"principal_type,omitempty"`
	PrincipalName  *string        `json:"principal_name,omitempty"`
	RealNameStatus *string        `json:"realname_status,omitempty"`
	WXVerifyInfo   *WXVerifyInfo  `json:"wx_verify_info,omitempty"`
	SignatureInfo  *SignatureInfo `json:"signature_info,omitempty"`
	HeadImageInfo  *HeadImageInfo `json:"head_image_info,omitempty"`
}

func (b BasicInfo) String() string {
	return Stringify(b)
}

// GetBasicInfo fetch account basic info.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html
func (s *AccountService) GetBasicInfo(ctx context.Context, token string) (*BasicInfo, *Response, error) {
	u := fmt.Sprintf("cgi-bin/account/getaccountbasicinfo?access_token=%v", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	info := new(BasicInfo)
	resp, err := s.client.Do(ctx, req, info)
	if err != nil {
		return nil, resp, err
	}
	return info, resp, nil
}
