package light

import (
	"github.com/sander6/hue/color"

	"bytes"
	"encoding/json"
	"testing"
)

func TestDeserializeSingleLightResponse(t *testing.T) {
	bulb := new(Bulb)
	err := json.NewDecoder(bytes.NewReader(singleLightResponse)).Decode(&bulb)
	if err != nil {
		t.Error(err)
		return
	}
	if bulb.Type != "Living Colors" {
		t.Errorf("wrong light type: %s", bulb.Type)
	}
	if bulb.Name != "LC 1" {
		t.Errorf("wrong name: %s", bulb.Name)
	}
	if bulb.ModelID != "LC0015" {
		t.Errorf("wrong model id: %s", bulb.ModelID)
	}
	if bulb.SWVersion != "1.0.3" {
		t.Errorf("wrong software version: %s", bulb.SWVersion)
	}
	if bulb.State.Position != On {
		t.Errorf("wrong on state: %v", bulb.State.Position)
	}
	if bulb.State.Hue != uint16(50000) {
		t.Errorf("wrong hue state: %d", bulb.State.Hue)
	}
	if bulb.State.Effect != NoEffect {
		t.Errorf("wrong effect state: %v", bulb.State.Effect)
	}
	if bulb.State.Alert != NoAlert {
		t.Errorf("wrong alert state: %v", bulb.State.Alert)
	}
	if bulb.State.Brightness != uint8(200) {
		t.Errorf("wrong brightness state: %d", bulb.State.Brightness)
	}
	if bulb.State.Saturation != uint8(200) {
		t.Errorf("wrong saturation state: %d", bulb.State.Saturation)
	}
	if bulb.State.ColorMode != color.HSMode {
		t.Errorf("wrong colormode state: %v", bulb.State.ColorMode)
	}
	if bulb.State.ColorTemp != color.Mirek(500) {
		t.Errorf("wrong color temperature state: %v", bulb.State.ColorTemp)
	}
	if bulb.State.Reachable != true {
		t.Errorf("wrong reachable state: %v", bulb.State.Reachable)
	}
	chroma, err := bulb.State.Chromaticity()
	if err != nil {
		t.Error(err)
		return
	}
	if chroma.X != 0.5 || chroma.Y != 0.5 {
		t.Errorf("wrong chromaticity state: %v", chroma)
	}
}

func TestDeserializeAllLightsResponse(t *testing.T) {
	bulbs := make(Bulbs, 50)
	err := json.NewDecoder(bytes.NewReader(allLightsResponse)).Decode(&bulbs)
	if err != nil {
		t.Error(err)
	}
	bulb, ok := bulbs["1"]
	if !ok {
		t.Error("missing light 1")
		return
	}
	if bulb.Type != "Extended color light" {
		t.Errorf("wrong light type: %s", bulb.Type)
	}
	if bulb.Name != "Hue Lamp 1" {
		t.Errorf("wrong name: %s", bulb.Name)
	}
	if bulb.ModelID != "LCT001" {
		t.Errorf("wrong model id: %s", bulb.ModelID)
	}
	if bulb.SWVersion != "66009461" {
		t.Errorf("wrong software version: %s", bulb.SWVersion)
	}
	if bulb.State.Position != On {
		t.Errorf("wrong on state: %v", bulb.State.Position)
	}
	if bulb.State.Hue != uint16(13088) {
		t.Errorf("wrong hue state: %d", bulb.State.Hue)
	}
	if bulb.State.Effect != NoEffect {
		t.Errorf("wrong effect state: %v", bulb.State.Effect)
	}
	if bulb.State.Alert != NoAlert {
		t.Errorf("wrong alert state: %v", bulb.State.Alert)
	}
	if bulb.State.Brightness != uint8(144) {
		t.Errorf("wrong brightness state: %d", bulb.State.Brightness)
	}
	if bulb.State.Saturation != uint8(212) {
		t.Errorf("wrong saturation state: %d", bulb.State.Saturation)
	}
	if bulb.State.ColorMode != color.XYMode {
		t.Errorf("wrong colormode state: %v", bulb.State.ColorMode)
	}
	if bulb.State.ColorTemp != color.Mirek(467) {
		t.Errorf("wrong color temperature state: %v", bulb.State.ColorTemp)
	}
	if bulb.State.Reachable != true {
		t.Errorf("wrong reachable state: %v", bulb.State.Reachable)
	}
	chroma, err := bulb.State.Chromaticity()
	if err != nil {
		t.Error(err)
		return
	}
	if chroma.X != 0.5128 || chroma.Y != 0.4147 {
		t.Errorf("wrong chromaticity state: %v", chroma)
	}

	bulb, ok = bulbs["2"]
	if !ok {
		t.Error("missing light 2")
		return
	}
	if bulb.Type != "Extended color light" {
		t.Errorf("wrong light type: %s", bulb.Type)
	}
	if bulb.Name != "Hue Lamp 2" {
		t.Errorf("wrong name: %s", bulb.Name)
	}
	if bulb.ModelID != "LCT001" {
		t.Errorf("wrong model id: %s", bulb.ModelID)
	}
	if bulb.SWVersion != "66009461" {
		t.Errorf("wrong software version: %s", bulb.SWVersion)
	}
	if bulb.State.Position != Off {
		t.Errorf("wrong on state: %v", bulb.State.Position)
	}
	if bulb.State.Hue != uint16(0) {
		t.Errorf("wrong hue state: %d", bulb.State.Hue)
	}
	if bulb.State.Effect != NoEffect {
		t.Errorf("wrong effect state: %v", bulb.State.Effect)
	}
	if bulb.State.Alert != NoAlert {
		t.Errorf("wrong alert state: %v", bulb.State.Alert)
	}
	if bulb.State.Brightness != uint8(0) {
		t.Errorf("wrong brightness state: %d", bulb.State.Brightness)
	}
	if bulb.State.Saturation != uint8(0) {
		t.Errorf("wrong saturation state: %d", bulb.State.Saturation)
	}
	if bulb.State.ColorMode != color.HSMode {
		t.Errorf("wrong colormode state: %v", bulb.State.ColorMode)
	}
	if bulb.State.ColorTemp != color.Mirek(0) {
		t.Errorf("wrong color temperature state: %v", bulb.State.ColorTemp)
	}
	if bulb.State.Reachable != true {
		t.Errorf("wrong reachable state: %v", bulb.State.Reachable)
	}
	chroma, err = bulb.State.Chromaticity()
	if err != nil {
		t.Error(err)
		return
	}
	if chroma.X != 0 || chroma.Y != 0 {
		t.Errorf("wrong chromaticity state: %v", chroma)
	}
}

var (
	allLightsResponse = []byte(`{
		"1": {
			"state": {
				"on": true,
				"bri": 144,
				"hue": 13088,
				"sat": 212,
				"xy": [0.5128,0.4147],
				"ct": 467,
				"alert": "none",
				"effect": "none",
				"colormode": "xy",
				"reachable": true
			},
			"type": "Extended color light",
			"name": "Hue Lamp 1",
			"modelid": "LCT001",
			"swversion": "66009461",
			"pointsymbol": {
				"1": "none",
				"2": "none",
				"3": "none",
				"4": "none",
				"5": "none",
				"6": "none",
				"7": "none",
				"8": "none"
			}
		},
		"2": {
			"state": {
				"on": false,
				"bri": 0,
				"hue": 0,
				"sat": 0,
				"xy": [0,0],
				"ct": 0,
				"alert": "none",
				"effect": "none",
				"colormode": "hs",
				"reachable": true
			},
			"type": "Extended color light",
			"name": "Hue Lamp 2",
			"modelid": "LCT001",
			"swversion": "66009461",
			"pointsymbol": {
				"1": "none",
				"2": "none",
				"3": "none",
				"4": "none",
				"5": "none",
				"6": "none",
				"7": "none",
				"8": "none"
			}
		}
	}`)

	singleLightResponse = []byte(`{
		"state": {
			"hue": 50000,
			"on": true,
			"effect": "none",
			"alert": "none",
			"bri": 200,
			"sat": 200,
			"ct": 500,
			"xy": [0.5, 0.5],
			"reachable": true,
			"colormode": "hs"
		},
		"type": "Living Colors",
		"name": "LC 1",
		"modelid": "LC0015",
		"swversion": "1.0.3",
		"pointsymbol": {
			"1": "none",
			"2": "none",
			"3": "none",
			"4": "none",
			"5": "none",
			"6": "none",
			"7": "none",
			"8": "none"
		}
	}`)
)
