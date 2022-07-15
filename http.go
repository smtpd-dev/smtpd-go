package smtpd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Credentials struct {
	Username string
	Password string
}

// HTTP struct
type HTTP struct {
	client *http.Client
}

// Get http call
func (h *HTTP) Get(endpoint string, headers map[string]string) ([]byte, int, error) {
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return do(h.client, req)
}

// Post http call
func (h *HTTP) Post(endpoint string, payload []byte, headers map[string]string) ([]byte, int, error) {
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(payload))
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return do(h.client, req)
}

// PostWithBasicAuth http call
func (h *HTTP) PostWithBasicAuth(endpoint string, credentials Credentials, payload []byte, headers map[string]string) ([]byte, int, error) {
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(payload))
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.SetBasicAuth(credentials.Username, credentials.Password)

	return do(h.client, req)
}

// Put http call
func (h *HTTP) Put(endpoint string, payload []byte, headers map[string]string) ([]byte, int, error) {
	req, _ := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(payload))
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return do(h.client, req)
}

// Delete http call
func (h *HTTP) Delete(endpoint string, headers map[string]string) ([]byte, int, error) {
	req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return do(h.client, req)
}

func do(c *http.Client, r *http.Request) ([]byte, int, error) {
	resp, err := c.Do(r)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, err
}

// BuildEndpointWithQueryParameters add parameters to endpoint URL
func (h *HTTP) BuildEndpointWithQueryParameters(endpoint string, parameters map[string]string) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	if len(parameters) >= 1 {
		q := u.Query()
		for k, v := range parameters {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}
