package light

import ()

type Bulb struct {
	ID        int    `json:-`
	Type      string `json:"type"`
	Name      string `json:"name"`
	ModelID   string `json:"modelid"`
	SWVersion string `json:"swversion"`
	State     State  `json:"state"`
}

func (b Bulb) Hue() uint16 {
	return b.State.Hue
}

func (b Bulb) Saturation() uint8 {
	return b.State.Saturation
}

func (b Bulb) Brightness() uint8 {
	return b.State.Brightness
}

type Bulbs map[string]*Bulb

type BulbIDs []string
