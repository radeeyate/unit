package unit

// https://en.wikipedia.org/wiki/Luminous_intensity
// https://en.wikipedia.org/wiki/Candela

// LuminousIntensity represents a SI unit for luminous intensity (in candela, cd)
type LuminousIntensity float64

// constants
const (
	Candela LuminousIntensity = 1e0
)

// Candela returns the luminous intensity in cd
func (l LuminousIntensity) Candela() float64 {
	return float64(l)
}
