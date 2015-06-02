package main

import (
	"fmt"
	"log"
	"net/http"
)

const urlPattern = "http://%s/api/%s/%s"

type Bridge struct {
	client     *http.Client
	host, user string
}

func NewBridge(host, user string) *Bridge {
	log.Printf("connecting to bridge at %s as %q\n", host, user)
	c := new(http.Client)
	c.Transport = loggingRoundTripper{http.DefaultTransport}
	return &Bridge{
		client: c,
		host:   host,
		user:   user,
	}
}

func (b Bridge) AllLights() *AllLightsController {
	return &AllLightsController{&b}
}

func (b Bridge) Light(id string) *LightController {
	return &LightController{&b, id}
}

func (b Bridge) AllGroups() *AllGroupsController {
	return &AllGroupsController{&b}
}

func (b Bridge) Group(id string) *GroupController {
	return &GroupController{&b, id}
}

func (b Bridge) urlFor(path string, args ...interface{}) string {
	return fmt.Sprintf(urlPattern, b.host, b.user, fmt.Sprintf(path, args...))
}
