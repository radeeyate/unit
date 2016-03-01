package unit

// LuminousFlux represents a SI unit for luminous flux (in lumen, lm)
type LuminousFlux float64

// constants
const (
	Lumen LuminousFlux = 1e0 // SI
)

// Lumen returns the luminous flux in lm
func (l LuminousFlux) Lumen() float64 {
	return float64(l)
}
