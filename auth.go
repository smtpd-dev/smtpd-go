package smtpd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OauthTokenResponse -
type OauthTokenResponse struct {
	TokenType    string   `json:"token_type"`
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
}

// RefreshTokenRequest is contains a RefreshToken used to revoke a RefreshToken OR refresh an AccessToken
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) Authenticate() error {
	if len(c.key) == 0 && len(c.secret) == 0 {
		return ErrInvalidCredentials
	}

	ok := c.validate()
	if ok {
		return nil
	}

	if c.hasRefreshToken() {
		return c.refreshAccessToken()
	}

	return c.basicAuth()
}

func (c *Client) basicAuth() error {
	body, statusCode, err := c.http.PostWithBasicAuth(
		fmt.Sprintf("%s/oauth/token?grant_type=password", baseURL),
		Credentials{
			c.key,
			c.secret,
		},
		nil,
		map[string]string{
			HeaderContentType: ContentTypeJson,
		},
	)
	if err != nil {
		return err
	}

	if statusCode != http.StatusOK {
		return c.parseAPIError(body)
	}

	var r OauthTokenResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	c.accessToken = r.AccessToken
	c.refreshToken = r.RefreshToken

	return nil
}

func (c *Client) validate() bool {
	if c.hasAccessToken() {
		_, statusCode, err := c.http.Get(
			fmt.Sprintf("%s/oauth/validate", baseURL),
			map[string]string{
				HeaderAuthorization: c.bearer(),
				HeaderContentType:   ContentTypeJson,
			})
		if err != nil {
			c.clearAccessToken()
			return false
		}

		if statusCode == http.StatusOK {
			return true
		}
	}
	return false
}

func (c *Client) refreshAccessToken() error {

	if !c.hasRefreshToken() {
		return c.basicAuth()
	}

	b, err := json.Marshal(RefreshTokenRequest{
		RefreshToken: c.refreshToken,
	})
	if err != nil {
		return err
	}

	body, statusCode, err := c.http.Post(
		fmt.Sprintf("%s/oauth/token?grant_type=refresh_token", baseURL),
		b,
		map[string]string{
			HeaderAuthorization: c.bearer(),
			HeaderContentType:   ContentTypeJson,
		})
	if err != nil {
		c.clearAccessToken()
		c.clearRefreshToken()
		return c.basicAuth()
	}

	if statusCode != http.StatusOK {
		return c.parseAPIError(body)
	}

	var r OauthTokenResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return err
	}

	c.accessToken = r.AccessToken

	return nil
}

func (c *Client) hasAccessToken() bool {
	return len(c.accessToken) != 0
}

func (c *Client) clearAccessToken() {
	c.accessToken = ""
}

func (c *Client) hasRefreshToken() bool {
	return len(c.refreshToken) != 0
}

func (c *Client) clearRefreshToken() {
	c.refreshToken = ""
}

func (c *Client) bearer() string {
	const (
		bearerPrefix string = "Bearer "
	)
	return bearerPrefix + c.accessToken
}
