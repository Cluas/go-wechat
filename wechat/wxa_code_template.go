package wechat

import (
	"context"
	"fmt"
)

// Draft represents a code template draft.
type Draft struct {
	CreateTime      *int    `json:"create_time,omitempty"`
	UserVersion     *string `json:"user_version,omitempty"`
	UserDescription *string `json:"user_desc,omitempty"`
	DraftID         *int    `json:"draft_id,omitempty"`
}

// TemplateDrafts represents a draft list.
type TemplateDrafts struct {
	DraftList []*Draft `json:"draft_list"`
}

// GetTemplateDrafts fetches template drafts.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatedraftlist.html
func (s *WXAService) GetTemplateDrafts(ctx context.Context, token string) (*TemplateDrafts, *Response, error) {
	u := fmt.Sprintf("wxa/gettemplatedraftlist?access_token=%v", token)
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
func (s *WXAService) AddDraftToTemplate(ctx context.Context, token string, draftID int) (*Response, error) {
	u := fmt.Sprintf("wxa/addtotemplate?access_token=%v", token)
	payload := &Draft{DraftID: Int(draftID)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}

// Template represents a code template.
type Template struct {
	CreateTime      *int    `json:"create_time,omitempty"`
	UserVersion     *string `json:"user_version,omitempty"`
	UserDescription *string `json:"user_desc,omitempty"`
	TemplateID      *int    `json:"template_id,omitempty"`
}

// Templates represents a template list.
type Templates struct {
	TemplateList []*Template `json:"template_list"`
}

// GetTemplates fetches template.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/gettemplatelist.html
func (s *WXAService) GetTemplates(ctx context.Context, token string) (*Templates, *Response, error) {
	u := fmt.Sprintf("wxa/gettemplatelist?access_token=%v", token)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	templates := new(Templates)
	resp, err := s.client.Do(ctx, req, templates)
	if err != nil {
		return nil, resp, err
	}
	return templates, resp, nil
}

// DeleteTemplateByID delete template by id.
//
// Wechat API docs:
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code_template/deletetemplate.html
func (s *WXAService) DeleteTemplateByID(ctx context.Context, token string, templateID int) (*Response, error) {
	u := fmt.Sprintf("wxa/deletetemplate?access_token=%v", token)
	payload := &Template{TemplateID: Int(templateID)}
	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
