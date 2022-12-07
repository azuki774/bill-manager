package factory

import (
	"azuki774/bill-manager/internal/repository"
)

type HTTPClientInfo struct {
	Scheme    string // http or https
	Host      string
	Port      string
	BasicAuth repository.BasicAuth
}

func setDefaultValueHTTPClient(opts *HTTPClientInfo) {
	if opts.Scheme == "" {
		opts.Scheme = "http"
	}
	if opts.Host == "" {
		opts.Host = "localhost"
	}
	if opts.Port == "" {
		opts.Port = "8888"
	}
	if opts.BasicAuth.User == "" {
		opts.BasicAuth.User = "test"
	}
	if opts.BasicAuth.Pass == "" {
		opts.BasicAuth.Pass = "test"
	}
}

func NewHTTPClient(opts *HTTPClientInfo) *repository.HTTPClient {
	setDefaultValueHTTPClient(opts)
	client := repository.HTTPClient{
		Scheme: opts.Scheme,
		Host:   opts.Host,
		Port:   opts.Port,
		BasicAuth: repository.BasicAuth{
			User: opts.BasicAuth.User,
			Pass: opts.BasicAuth.Pass,
		},
	}
	return &client
}
