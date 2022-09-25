package repository

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", "https://example.com:443/api/record",
		httpmock.NewStringResponder(201, "created"),
	)
	m.Run()
}

func TestHTTPClient_PostJson(t *testing.T) {
	type fields struct {
		Scheme    string
		Host      string
		Port      string
		BasicAuth BasicAuth
	}
	type args struct {
		ctx      context.Context
		endPoint string
		reqBody  []byte
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantResBody    []byte
		wantStatusCode int
		wantErr        bool
	}{
		{
			name: "ok",
			fields: fields{
				Scheme: "https",
				Host:   "example.com",
				Port:   "443",
				BasicAuth: BasicAuth{
					User: "user",
					Pass: "pass",
				},
			},
			args: args{
				ctx:      context.Background(),
				endPoint: "/api/record",
				reqBody:  []byte("test"),
			},
			wantResBody:    []byte("created"),
			wantStatusCode: http.StatusCreated,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPClient{
				Scheme:    tt.fields.Scheme,
				Host:      tt.fields.Host,
				Port:      tt.fields.Port,
				BasicAuth: tt.fields.BasicAuth,
			}
			gotResBody, gotStatusCode, err := c.PostJson(tt.args.ctx, tt.args.endPoint, tt.args.reqBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPClient.PostJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResBody, tt.wantResBody) {
				t.Errorf("HTTPClient.PostJson() gotResBody = %v, want %v", gotResBody, tt.wantResBody)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("HTTPClient.PostJson() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
