package unit

// https://en.wikipedia.org/wiki/Illuminance
// https://en.wikipedia.org/wiki/Lux

// Illuminance represents a SI unit for illuminance (in lux, lx)
type Illuminance float64

// constants
const (
	Lux Illuminance = 1e0 // SI
)

// Lux returns the illuminance in lx
func (i Illuminance) Lux() float64 {
	return float64(i)
}
