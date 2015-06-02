package light

import ()

type Position bool

const (
	Off Position = false
	On           = true
)

func (p Position) String() string {
	if bool(p) {
		return "ON"
	}
	return "OFF"
}
