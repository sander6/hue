package main

import (
	"github.com/sander6/hue/light"

	"fmt"
	"io"
	"strings"
)

type (
	LightPrinter struct {
		*light.Bulb
	}

	GroupPrinter struct {
		*light.Group
	}
)

func (lp LightPrinter) WriteTo(w io.Writer) (int, error) {
	var x, y float64
	chr, err := lp.Bulb.State.Chromaticity()
	if err != nil {
		x, y = 0, 0
	} else {
		x, y = chr.X, chr.Y
	}
	return io.WriteString(w, fmt.Sprintf("%2d\t%-32s\t%3s\t%-0.4f\t%-0.4f\n", lp.Bulb.ID(), lp.Bulb.Name, lp.Bulb.State.Position.String(), x, y))
}

func (gp GroupPrinter) WriteTo(w io.Writer) (int, error) {
	return io.WriteString(w, fmt.Sprintf("%2d\t%-32s\t%s\n", gp.Group.ID(), gp.Group.Name, strings.Join(gp.Group.Lights, ", ")))
}
