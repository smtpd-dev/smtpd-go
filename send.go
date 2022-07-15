package smtpd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendBasicMessage is used to send an email
type SendBasicMessage struct {
	RecipientEmailAddress string   `json:"recipient_email_address"`
	RecipientName         string   `json:"recipient_name"`
	FromEmailAddress      string   `json:"from_email_address"`
	FromName              string   `json:"from_name"`
	ReplyTo               string   `json:"reply_to"`
	Subject               string   `json:"subject"`
	ContentHTML           string   `json:"content_html"`
	ContentText           string   `json:"content_text"`
	OpenTracking          bool     `json:"open_tracking"`
	ClickTracking         bool     `json:"click_tracking"`
	Tags                  []string `json:"tags"`
}

// SendBasicDetailedResponse after a message has been accepted
type SendBasicDetailedResponse struct {
	SendID                string   `json:"send_id"`
	SendKey               string   `json:"send_key"`
	SendingProfileID      string   `json:"sending_profile_id"`
	FromEmailAddress      string   `json:"from_email_address"`
	FromName              string   `json:"from_name"`
	RecipientEmailAddress string   `json:"recipient_email_address"`
	RecipientName         string   `json:"recipient_name"`
	Subject               string   `json:"subject"`
	OpenTracking          bool     `json:"open_tracking"`
	ClickTracking         bool     `json:"click_tracking"`
	Tags                  []string `json:"tags"`
	State                 string   `json:"state"`
	StateCategory         string   `json:"state_category"`
	Error                 string   `json:"error,omitempty"`
	CreatedAtUtc          int64    `json:"created_at_utc"`
	ModifiedAtUtc         int64    `json:"modified_at_utc"`
}

// SendBasic sends a transactional email
func (c *Client) SendBasic(message SendBasicMessage) (SendBasicDetailedResponse, error) {
	var result SendBasicDetailedResponse

	if !c.preFlight() {
		return result, ErrInvalidCredentials
	}

	b, err := json.Marshal(&message)
	if err != nil {
		return result, err
	}

	body, statusCode, err := c.http.Post(
		fmt.Sprintf("%s/api/%s/email/send", baseURL, baseVersion),
		b,
		c.baseHeaders(),
	)
	if err != nil {
		return result, err
	}

	if statusCode != http.StatusAccepted {
		return result, c.parseAPIError(body)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
