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
	id     int        `json:-`
	Type   *GroupType `json:type`
	Name   string     `json:name`
	Lights BulbIDs    `json:lights`
	State  State      `json:action`
}

func (g Group) ID() int {
	return g.id
}

func (g Group) Hue() uint16 {
	return g.State.Hue
}

func (g Group) Saturation() uint8 {
	return g.State.Saturation
}

func (g Group) Brightness() uint8 {
	return g.State.Brightness
}

type GroupType string

const (
	LuminaireType   GroupType = "Luminaire"
	LightsourceType           = "Lightsource"
	GenericType               = "LightGroup"
)

type Groups []*Group

func (g *Groups) UnmarshalJSON(body []byte) error {
	gmap := make(map[string]*Group)
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&gmap)
	if err != nil {
		return err
	}
	gs := make(Groups, len(gmap))
	i := 0
	for id, group := range gmap {
		group.id, err = strconv.Atoi(id)
		if err != nil {
			return err
		}
		gs[i] = group
		i++
	}
	*g = gs
	return nil
}
