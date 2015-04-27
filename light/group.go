package light

import (
	"bytes"
	"encoding/json"
	"strconv"
)

var MaxSupportedGroups = 100

// Group represents one or more bulbs that receive actions together
// Currently you can register up to 100 Groups with a Hue bridge
type Group struct {
	ID     int        `json:-`
	Type   *GroupType `json:type`
	Name   string     `json:name`
	Lights BulbIDs    `json:lights`
	Action State      `json:action`
}

func (g Group) Hue() uint16 {
	return g.Action.Hue
}

func (g Group) Saturation() uint8 {
	return g.Action.Saturation
}

func (g Group) Brightness() uint8 {
	return g.Action.Brightness
}

type GroupType string

const (
	LuminaireType   GroupType = "Luminaire"
	LightsourceType           = "Lightsource"
	GenericType               = "LightGroup"
)

type Groups map[int]*Group

func (gs Groups) UnmarshalJSON(body []byte) error {
	var m map[string]*Group
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&m)
	if err != nil {
		return err
	}
	for key, group := range m {
		idx, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		group.ID = idx
		gs[idx] = group
	}
	return nil
}
