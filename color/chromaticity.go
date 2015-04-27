package color

import (
	"bytes"
	"encoding/json"
	"log"
)

func AtXY(x, y float64) *Chromaticity {
	return &Chromaticity{X: x, Y: y}
}

// Chromaticity represents an x, y point on the chromaticity plane
type Chromaticity struct {
	X, Y float64
}

func (c *Chromaticity) UnmarshalJSON(j []byte) error {
	var vec [2]float64
	err := json.NewDecoder(bytes.NewReader(j)).Decode(&vec)
	if err != nil {
		log.Printf("decode error for Chromaticity on %s\n", j)
		return err
	}
	c.X = vec[0]
	c.Y = vec[1]
	return nil
}
