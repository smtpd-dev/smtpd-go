package smtpd

import (
	"net/http"
	"testing"
)

func TestClient_parseAPIError(t *testing.T) {
	type fields struct {
		http         HTTP
		key          string
		secret       string
		accessToken  string
		refreshToken string
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Fail/EmptyBody",
			fields: fields{
				http: HTTP{
					client: &http.Client{},
				},
				key:          "TEST",
				secret:       "TEST",
				accessToken:  "",
				refreshToken: "",
			},
			args: args{
				body: nil,
			},
			wantErr: true,
		},
		{
			name: "Pass/SimpleErrorPayload",
			fields: fields{
				http: HTTP{
					client: &http.Client{},
				},
				key:          "TEST",
				secret:       "TEST",
				accessToken:  "",
				refreshToken: "",
			},
			args: args{
				body: []byte(`{"code": 40106,"message": "Refresh token is invalid"}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http:         tt.fields.http,
				key:          tt.fields.key,
				secret:       tt.fields.secret,
				accessToken:  tt.fields.accessToken,
				refreshToken: tt.fields.refreshToken,
			}
			if err := c.parseAPIError(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Client.parseAPIError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_apiError_String(t *testing.T) {
	type fields struct {
		Code    int
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Pass/SimpleError",
			fields: fields{
				Code:    40410,
				Message: "resource not found",
			},
			want: "code: 40410, message: resource not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &apiError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("apiError.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_apiError_Error(t *testing.T) {
	type fields struct {
		Code    int
		Message string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Pass/SimpleError",
			fields: fields{
				Code:    40410,
				Message: "resource not found",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &apiError{
				Code:    tt.fields.Code,
				Message: tt.fields.Message,
			}
			if err := e.Error(); (err != nil) != tt.wantErr {
				t.Errorf("apiError.Error() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
