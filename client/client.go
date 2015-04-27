// package client provides an HTTP client for communicating with a Philips Hue Bridge
package client

import (
	"log"
	"net/http"
)

type Client struct {
	client     *http.Client
	host, user string
}

func New(host, user string) *Client {
	log.Printf("connecting to bridge at %s as %q\n", host, user)
	return &Client{
		client: http.DefaultClient,
		host:   host,
		user:   user,
	}
}
