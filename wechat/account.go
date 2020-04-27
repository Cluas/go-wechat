package wechat

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// AccountService Wechat API docs: https://developers.weixin.qq.com/doc/
type AccountService service

// WXVerifyInfo represents wechat verify info.
type WXVerifyInfo struct {
	QualificationVerify   bool       `json:"qualification_verify"`
	NamingVerify          bool       `json:"naming_verify"`
	AnnualReview          bool       `json:"annual_review"`
	AnnualReviewBeginTime *time.Time `json:"annual_review_begin_time"`
	AnnualReviewEndTime   *time.Time `json:"annual_review_end_time"`
}

// SignatureInfo represents wechat signature info.
type SignatureInfo struct {
	Signature       string `json:"signature"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

// HeadImageInfo represents wechat head image info.
type HeadImageInfo struct {
	HeadImageURL    string `json:"head_image_url"`
	ModifyUsedCount int    `json:"modify_used_count"`
	ModifyQuota     int    `json:"modify_quota"`
}

// AccountBasicInfo represents wechat basic info.
type AccountBasicInfo struct {
	AppID          string         `json:"app_id"`
	AccountType    int            `json:"account_type"`
	PrincipalType  int            `json:"principal_type"`
	PrincipalName  string         `json:"principal_name"`
	RealNameStatus string         `json:"realname_status"`
	WXVerifyInfo   *WXVerifyInfo  `json:"wx_verify_info"`
	SignatureInfo  *SignatureInfo `json:"signature_info"`
	HeadImageInfo  *HeadImageInfo `json:"head_image_info"`
}

// GetAccountBasicInfo fetch account basic info.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html
func (s *AccountService) GetAccountBasicInfo(ctx context.Context, token string) (*AccountBasicInfo, *Response, error) {
	u := fmt.Sprintf("cgi-bin/account/getaccountbasicinfo?access_token=%v", token)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	info := new(AccountBasicInfo)
	resp, err := s.client.Do(ctx, req, info)
	if err != nil {
		return nil, resp, err
	}
	return info, resp, nil
}
