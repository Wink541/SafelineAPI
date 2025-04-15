package safeLineApi

import (
	"net/url"
	"strconv"
)

func (u *URL) SSLCertUrl() string {
	path := "/api/open/cert"
	u.Path = path
	return (*url.URL)(u).String()
}

func (u *URL) SSLCertUrlWithParam(id int) string {
	path := "/api/open/cert/" + strconv.Itoa(id)
	u.Path = path
	return (*url.URL)(u).String()
}
