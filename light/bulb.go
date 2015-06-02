package light

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type Bulb struct {
	id        int    `json:-`
	Type      string `json:"type"`
	Name      string `json:"name"`
	ModelID   string `json:"modelid"`
	SWVersion string `json:"swversion"`
	State     State  `json:"state"`
}

func (b Bulb) ID() int {
	return b.id
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

type Bulbs []*Bulb

func (b *Bulbs) UnmarshalJSON(body []byte) error {
	bmap := make(map[string]*Bulb)
	err := json.NewDecoder(bytes.NewReader(body)).Decode(&bmap)
	if err != nil {
		return err
	}
	bs := make(Bulbs, len(bmap))
	i := 0
	for id, bulb := range bmap {
		bulb.id, err = strconv.Atoi(id)
		if err != nil {
			return err
		}
		bs[i] = bulb
		i++
	}
	*b = bs
	return nil
}

type BulbIDs []string

func (b BulbIDs) Len() int {
	return len(b)
}

func (b BulbIDs) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b BulbIDs) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
