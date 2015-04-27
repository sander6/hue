package client

import (
	"github.com/sander6/hue/light"

	"encoding/json"
	"fmt"
)

var (
	MaxSupportedLights int = 50
)

func (c *Client) FindAllLights() (light.Bulbs, error) {
	resp, err := c.client.Get(fmt.Sprintf("http://%s/api/%s/lights", c.host, c.user))
	if err != nil {
		return nil, err
	}
	bulbs := make(light.Bulbs, MaxSupportedLights)
	err = json.NewDecoder(resp.Body).Decode(&bulbs)
	if err != nil {
		return nil, err
	}
	return bulbs, nil
}

func (c *Client) FindLight(id string) (*light.Bulb, error) {
	resp, err := c.client.Get(fmt.Sprintf("http://%s/api/%s/lights/%s", c.host, c.user, id))
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

func (c *Client) SearchForLights() error {
	return nil
}

func (c *Client) GetNewLights() (light.Bulbs, error) {
	return nil, nil
}
