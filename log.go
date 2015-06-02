package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

type loggingRoundTripper struct {
	http.RoundTripper
}

func (rt loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	body, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Println(string(body))
	}
	resp, err := rt.RoundTripper.RoundTrip(req)
	if err == nil {
		body, err = httputil.DumpResponse(resp, true)
		if err == nil {
			log.Println(string(body))
		}
	}
	return resp, err
}
