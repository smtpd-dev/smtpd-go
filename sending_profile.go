package smtpd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateSendingProfile represents the request payload to create a sending profile.
type CreateSendingProfile struct {
	ProfileName               string `json:"profile_name"`
	SendingDomain             string `json:"sending_domain"`
	LinkDomain                string `json:"link_domain"`
	LinkDomainDefaultRedirect string `json:"link_domain_default_redirect"`
	BounceDomain              string `json:"bounce_domain"`
}

// SendingProfile is return when doing a lookup
type SendingProfile struct {
	SendingProfileID          string `json:"profile_id"`
	ProfileName               string `json:"profile_name"`
	SendingDomain             string `json:"sending_domain"`
	LinkDomain                string `json:"link_domain"`
	LinkDomainDefaultRedirect string `json:"link_domain_default_redirect"`
	BounceDomain              string `json:"bounce_domain"`
	State                     string `json:"state"`
	CreatedAtUtc              int64  `json:"created_at_utc"`
	ModifiedAtUtc             int64  `json:"modified_at_utc"`
}

// SendingProfileSetup is return when doing a lookup
type SendingProfileSetup struct {
	SendingProfileID string          `json:"profile_id"`
	ProfileName      string          `json:"profile_name"`
	Records          []DomainRecords `json:"dns_records"`
	State            string          `json:"state"`
	CreatedAtUtc     int64           `json:"created_at_utc"`
	ModifiedAtUtc    int64           `json:"modified_at_utc"`
}

// DomainRecords is the DKIM records for a domain
type DomainRecords struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	ValidationState string `json:"validation_state"`
}

// CreateProfile creates a profile
func (c *Client) CreateProfile(profile *CreateSendingProfile) (SendingProfile, error) {
	var result SendingProfile

	if !c.preFlight() {
		return result, ErrInvalidCredentials
	}

	b, err := json.Marshal(profile)
	if err != nil {
		return result, err
	}

	body, statusCode, err := c.http.Post(
		fmt.Sprintf("%s/api/%s/email/profile", baseURL, baseVersion),
		b,
		c.baseHeaders(),
	)
	if err != nil {
		return result, err
	}

	if statusCode != http.StatusAccepted {
		return result, c.parseResponseError(body)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetProfileByID retrieves an email sending profile by ID
func (c *Client) GetProfileByID(id string) (SendingProfile, error) {
	var result SendingProfile

	if !c.preFlight() {
		return result, ErrInvalidCredentials
	}

	body, statusCode, err := c.http.Get(
		fmt.Sprintf("%s/api/%s/email/profile/%s/setup", baseURL, baseVersion, id),
		c.baseHeaders(),
	)
	if err != nil {
		return result, err
	}

	if statusCode != http.StatusOK {
		return result, c.parseResponseError(body)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetProfileByName retrieves an email sending profile by name
func (c *Client) GetProfileByName(name string) (SendingProfile, error) {
	var result SendingProfile

	if !c.preFlight() {
		return result, ErrInvalidCredentials
	}

	body, statusCode, err := c.http.Get(
		fmt.Sprintf("%s/api/%s/email/profile", baseURL, baseVersion),
		c.baseHeaders(),
	)
	if err != nil {
		return result, err
	}

	if statusCode != http.StatusAccepted {
		return result, c.parseResponseError(body)
	}

	var profiles []SendingProfile
	err = json.Unmarshal(body, &profiles)
	if err != nil {
		return result, err
	}

	for _, profile := range profiles {
		if profile.ProfileName == name {
			return profile, nil
		}
	}
	return result, ErrResourceNotFound
}

// GetAllProfiles retrieves all email sending profiles
func (c *Client) GetAllProfiles() ([]SendingProfile, error) {
	var result []SendingProfile

	if !c.preFlight() {
		return result, ErrInvalidCredentials
	}

	body, statusCode, err := c.http.Get(
		fmt.Sprintf("%s/api/%s/email/profile", baseURL, baseVersion),
		c.baseHeaders(),
	)
	if err != nil {
		return result, err
	}

	if statusCode != http.StatusOK {
		return result, c.parseResponseError(body)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
