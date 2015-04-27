package client

import (
	"github.com/sander6/hue/light"
)

func (c *Client) FindAllGroups() (light.Groups, error) {
	return nil, nil
}

func (c *Client) FindGroup(id int) (*light.Group, error) {
	return nil, nil
}

func (c *Client) CreateNewGroup(bulbs ...*light.Bulb) (*light.Group, error) {
	return nil, nil
}
