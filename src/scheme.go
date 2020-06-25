package main

import "net/url"

const (
	SchemeHTTP string = "http"

	SchemeHTTPS string = "https"
)

func isVaildScheme(uri string) bool {

	if u, _ := url.Parse(uri); u.Scheme == SchemeHTTP || u.Scheme == SchemeHTTPS {
		return true
	}

	return false
}

func getScheme(uri string) string {
	u, _ := url.Parse(uri)

	switch u.Scheme {
	case SchemeHTTP:
		return SchemeHTTP
	case SchemeHTTPS:
		return SchemeHTTPS
	default:
		return ""
	}
}
