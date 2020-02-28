package utils

import (
	"net/http"
	"net/http/cookiejar"
)

func CreateClientHttp() *http.Client {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	return client
}
