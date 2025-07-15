package smtpd

import (
	"net/http"
	"time"
)

const (
	baseURL     string = "https://api.smtpd.dev" // baseURL - Base URL for the SMTPD API
	baseVersion string = "v1"                    // baseVersion Base version for the API
	client      string = "smtpd-go/0.1"
)

// Client -
type Client struct {
	http         HTTP
	key          string
	secret       string
	accessToken  string
	refreshToken string
}

// New returns a new SMTPD instance
func New(key, secret string) *Client {
	return &Client{
		http: HTTP{
			client: &http.Client{
				Timeout: 10 * time.Second,
				Transport: &http.Transport{
					TLSHandshakeTimeout:   10 * time.Second,
					ResponseHeaderTimeout: 10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				},
			},
		},
		key:    key,
		secret: secret,
	}
}

func (c *Client) preFlight() bool {
	if !c.hasAccessToken() {
		err := c.Authenticate()
		if err != nil {
			return false
		}
	}

	return true
}

func (c *Client) baseHeaders() map[string]string {
	return map[string]string{
		HeaderUserAgent:     client,
		HeaderContentType:   ContentTypeJson,
		HeaderAuthorization: c.bearer(),
	}
}
