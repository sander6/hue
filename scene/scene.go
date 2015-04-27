// package scene provides models for scenes in a Hue system
package scene

import (
	"github.com/sander6/hue/light"
)

type Scene struct {
	Name   string        `json:"name"`
	Active bool          `json:"active"`
	Lights light.BulbIDs `json:"lights"`
}

type Scenes map[string]*Scene
