package esms

import (
	"net/http"
	"net/url"
	"time"
)

// BalanceResponse represents balance response of the current account
type BalanceResponse struct {
	Balance      int64      `json:"Balance"`
	CodeResponse StatusCode `json:"CodeResponse"`
	UserID       int        `json:"UserID"`
}

// SendMultipleMessageInput represents message request input
type SendMultipleMessageInput struct {
	Credential
	// message content
	Content string `json:"Content"`
	// phone number of the recipient
	Phone       string      `json:"Phone"`
	IsUnicode   string      `json:"IsUnicode,omitempty"`
	BrandName   string      `json:"Brandname,omitempty"`
	SmsType     SmsType     `json:"SmsType"`
	Sandbox     SandboxMode `json:"Sandbox,omitempty"`
	RequestId   string      `json:"RequestId,omitempty"`
	CallbackUrl string      `json:"CallbackUrl,omitempty"`
	CampaignID  string      `json:"campaignid,omitempty"`
	SendDate    time.Time   `json:"SendDate,omitempty"`
}

// SendMultipleMessageResponse represents send multiple message response
type SendMultipleMessageResponse struct {
	CodeResult      StatusCode `json:"CodeResult"`
	CountRegenerate int        `json:"CountRegenerate"`
	SmsID           string     `json:"SMSID"`
	ErrorMessage    string     `json:"ErrorMessage,omitempty"`
}

type smsService struct {
	client *httpClient
}

// GetBalance request balance information of the current account
func (ss *smsService) GetBalance() (*BalanceResponse, *http.Response, error) {
	u, err := url.Parse("/GetBalance_json")
	if err != nil {
		return nil, nil, err
	}
	// create the request
	req, err := ss.client.NewRequest("POST", u.String(), ss.client.credential)
	if err != nil {
		return nil, nil, err
	}

	result := &BalanceResponse{}
	resp, err := ss.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}

// SendMultipleMessage send multiple SMS message to target phone numbers
func (ss *smsService) SendMultipleMessage(input SendMultipleMessageInput) (*SendMultipleMessageResponse, *http.Response, error) {
	u, err := url.Parse("/SendMultipleMessage_V4_post_json")
	if err != nil {
		return nil, nil, err
	}
	// create the request
	input.Credential = ss.client.credential
	req, err := ss.client.NewRequest("POST", u.String(), input)
	if err != nil {
		return nil, nil, err
	}

	result := &SendMultipleMessageResponse{}
	resp, err := ss.client.Do(req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, err
}
