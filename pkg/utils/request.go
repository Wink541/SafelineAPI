package utils

import (
	"crypto/tls"
	"io"
	"net/http"
)

func Request(method, urlStr string, body io.Reader, header http.Header) (*http.Response, error) {
	req, _ := http.NewRequest(method, urlStr, body)
	req.Header = header
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
