package smtpd

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// ErrInvalidCredentials is return when the api key & secret are invalid
	ErrInvalidCredentials = errors.New("api key OR secret is invalid")

	// ErrResourceNotFound is return when the resource is not found
	ErrResourceNotFound = errors.New("resource not found")
)

func (c *Client) parseAPIError(body []byte) error {
	var errCode apiError

	err := json.Unmarshal(body, &errCode)
	if err != nil {
		return err
	}

	return errCode.Error()
}

//smtpdError is a error response
type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *apiError) String() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *apiError) Error() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}
