package light

import (
	"github.com/sander6/hue/color"

	"encoding/json"
)

type State struct {
	Position   Position        `json:"on"`
	Hue        uint16          `json:"hue"`
	Saturation uint8           `json:"sat"`
	Brightness uint8           `json:"bri"`
	ColorMode  color.Mode      `json:"colormode"`
	ColorTemp  color.Mirek     `json:"ct"`
	XY         json.RawMessage `json:"xy"`
	Effect     Effect          `json:"effect"`
	Alert      Alert           `json:"alert"`
	Reachable  bool            `json:"reachable"`
}

func (s State) Chromaticity() (*color.Chromaticity, error) {
	var arr [2]float64
	err := json.Unmarshal(s.XY, &arr)
	if err != nil {
		return nil, err
	}
	chroma := new(color.Chromaticity)
	chroma.X = arr[0]
	chroma.Y = arr[1]
	return chroma, nil
}
