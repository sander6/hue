// package color provides numerous representations of colors
package color

import (
	"errors"
)

var ErrDecode = errors.New("decode error")

// Color contains hue, saturation, and brightness values corresponding to some unique color.
type Color interface {
	Hue() uint16
	Saturation() uint8
	Brightness() uint8
}

// ToRGB is a helper function to turn a Color into RGB (for easy display on the web)
func ToRGB(c Color) *RGB {
	return &RGB{}
}

type Mode string

const (
	HSMode Mode = "hs"
	XYMode      = "xy"
	CTMode      = "ct"
)
