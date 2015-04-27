package color

// Color temperature represents white light. Warmer temperatures are
// more yellow, which colder temperatures are more blue.
type Temperature interface {
	Mirek() int
}

// Kelvin represents a color temperature in degrees kelvin
type Kelvin int

// Mirek represents a color temperature in reciprocal megakelvins
type Mirek int
