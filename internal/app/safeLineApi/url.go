package safeLineApi

import "net/url"

type URL url.URL

func (u *URL) String() string {
	return (*url.URL)(u).String()
}
