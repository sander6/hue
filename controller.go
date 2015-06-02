package main

import (
	"github.com/sander6/hue/light"

	"bytes"
	"encoding/json"
	"net/http"
)

type (
	AllLightsController struct {
		bridge *Bridge
	}

	LightController struct {
		bridge *Bridge
		id     string
	}

	AllGroupsController struct {
		bridge *Bridge
	}

	GroupController struct {
		bridge *Bridge
		id     string
	}
)

func (c AllLightsController) Find() (light.Bulbs, error) {
	resp, err := c.bridge.client.Get(c.bridge.urlFor("lights"))
	if err != nil {
		return nil, err
	}
	var bulbs light.Bulbs
	err = json.NewDecoder(resp.Body).Decode(&bulbs)
	if err != nil {
		return nil, err
	}
	return bulbs, nil
}

func (c AllLightsController) New() error {
	return nil
}

func (c AllLightsController) Search() (light.Bulbs, error) {
	return nil, nil
}

func (c AllLightsController) On() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": true})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("groups/%s/action", "0"), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c AllLightsController) Off() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": false})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("groups/%s/action", "0"), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c LightController) Find() (*light.Bulb, error) {
	resp, err := c.bridge.client.Get(c.bridge.urlFor("lights/%s", c.id))
	if err != nil {
		return nil, err
	}
	bulb := new(light.Bulb)
	err = json.NewDecoder(resp.Body).Decode(&bulb)
	if err != nil {
		return nil, err
	}
	return bulb, nil
}

func (c LightController) On() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": true})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("lights/%s/state", c.id), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c LightController) Off() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": false})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("lights/%s/state", c.id), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c LightController) Alert() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]string{"alert": "select"})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("lights/%s/state", c.id), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c AllGroupsController) Find() (light.Groups, error) {
	resp, err := c.bridge.client.Get(c.bridge.urlFor("groups"))
	if err != nil {
		return nil, err
	}
	var groups light.Groups
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (c GroupController) On() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": true})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("groups/%s/action", c.id), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}

func (c GroupController) Off() error {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(map[string]bool{"on": false})
	req, err := http.NewRequest("PUT", c.bridge.urlFor("groups/%s/action", c.id), buf)
	if err != nil {
		return err
	}
	_, err = c.bridge.client.Do(req)
	return err
}
