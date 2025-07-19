package safeLineApi

import (
	"net/url"
)

func (u *URL) AuthTokenUrl() string {
	path := "/api/open/auth/token"
	u.Path = path
	return (*url.URL)(u).String()
}
