package wechat

import (
	"context"
	"fmt"
	"time"
)

// Draft docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
// #%E8%8D%89%E7%A8%BF%E4%BF%A1%E6%81%AF%E8%AF%B4%E6%98%8E
type Draft struct {
	CreateTime      *time.Time `json:"create_time,omitempty"`
	UserVersion     *string    `json:"user_version,omitempty"`
	UserDescription *string    `json:"user_desc,omitempty"`
	DraftID         *int       `json:"draft_id,omitempty"`
}

// TemplateDrafts docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type TemplateDrafts struct {
	DraftList []*Draft `json:"draft_list"`
}

// GetTemplateDrafts fetches template drafts.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (s *WXAService) GetTemplateDrafts(ctx context.Context, token string) (*TemplateDrafts, *Response, error) {
	u := fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%s", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	drafts := new(TemplateDrafts)
	resp, err := s.client.Do(ctx, req, drafts)
	if err != nil {
		return nil, resp, err
	}
	return drafts, resp, nil
}

// AddDraftToTemplate add draft to template.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/addtotemplate.html
func (s *service) AddDraftToTemplate(ctx context.Context, token string, draftID int) (*Response, error) {
	u := fmt.Sprintf("wxa/addtotemplate?access_token=%s", token)
	payload := &Draft{DraftID: Int(draftID)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Template docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
// #%E6%A8%A1%E6%9D%BF%E4%BF%A1%E6%81%AF%E8%AF%B4%E6%98%8E
type Template struct {
	CreateTime      *time.Time `json:"create_time,omitempty"`
	UserVersion     *string    `json:"user_version,omitempty"`
	UserDescription *string    `json:"user_desc,omitempty"`
	TemplateID      *int       `json:"template_id,omitempty"`
}

// Templates docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
// #%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E
type Templates struct {
	TemplateList []*Template `json:"template_list"`
}

// GetTemplates fetches template.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (s *WXAService) GetTemplates(ctx context.Context, token string) (*TemplateDrafts, *Response, error) {
	u := fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%s", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	drafts := new(TemplateDrafts)
	resp, err := s.client.Do(ctx, req, drafts)
	if err != nil {
		return nil, resp, err
	}
	return drafts, resp, nil
}

// DeleteTemplateByID delete template by id.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (s *WXAService) DeleteTemplateByID(ctx context.Context, token string, templateID int) (*Response, error) {
	u := fmt.Sprintf("wxa/deletetemplate?access_token=%s", token)
	payload := &Template{TemplateID: Int(templateID)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
