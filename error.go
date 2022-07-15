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

func (c *Client) parseResponseError(body []byte) error {
	var errCode Error

	err := json.Unmarshal(body, &errCode)
	if err != nil {
		return err
	}

	return fmt.Errorf("code: %d, message: %s", errCode.Code, errCode.Message)
}

//Error Code response struct.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
