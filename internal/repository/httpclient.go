package repository

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
)

type HTTPClient struct {
	Scheme    string // http or https
	Host      string
	Port      string
	BasicAuth BasicAuth
}

type BasicAuth struct {
	User string
	Pass string
}

func (c *HTTPClient) PostJson(ctx context.Context, endPoint string, reqBody []byte) (resBody []byte, statusCode int, err error) {
	url := c.Scheme + "://" + net.JoinHostPort(c.Host, c.Port) + endPoint

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.BasicAuth.User, c.BasicAuth.Pass)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, err
	}
	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	return resBody, res.StatusCode, nil
}
