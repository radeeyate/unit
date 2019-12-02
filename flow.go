package unit

// Flow represents a SI unit of volume flow rate in cubic meters per second, m3/s
type Flow Unit

// ...
const (
	// SI
	CubicMeterPerSecond Flow = 1e0
	LiterPerSecond           = CubicMeterPerSecond * 1e-3
	LiterPerMinute           = LiterPerSecond * 1 / 60
	LiterPerHour             = LiterPerSecond * 1 / 3600

	// Non-SI
	CubicFootPerSecond = CubicMeterPerSecond * 0.028316846592

	// US liquid
	USLiquidGallonPerSecond = LiterPerSecond * 3.7854117839988
	USLiquidGallonPerMinute = USLiquidGallonPerSecond * 1 / 60
	USLiquidGallonPerHour   = USLiquidGallonPerSecond * 1 / 3600
)

// CubicMetersPerSecond returns the flow rate in m3/s
func (f Flow) CubicMetersPerSecond() float64 {
	return float64(f)
}

// CubicFeetPerSecond returns the flow rate in ft3/s
func (f Flow) CubicFeetPerSecond() float64 {
	return float64(f / CubicFootPerSecond)
}

// LitersPerSecond returns the flow rate in l/s
func (f Flow) LitersPerSecond() float64 {
	return float64(f / LiterPerSecond)
}

// LitersPerHour returns the flow rate in l/h
func (f Flow) LitersPerHour() float64 {
	return float64(f / LiterPerHour)
}

// LitersPerMinute returns the flow rate in l/min
func (f Flow) LitersPerMinute() float64 {
	return float64(f / LiterPerMinute)
}

// USLiquidGallonsPerHour returns the flow rate in gal/h
func (f Flow) USLiquidGallonsPerHour() float64 {
	return float64(f / USLiquidGallonPerHour)
}

// USLiquidGallonsPerMinute returns the flow rate in gal/min
func (f Flow) USLiquidGallonsPerMinute() float64 {
	return float64(f / USLiquidGallonPerMinute)
}

// USLiquidGallonsPerSecond returns the flow rate in gal/s
func (f Flow) USLiquidGallonsPerSecond() float64 {
	return float64(f / USLiquidGallonPerSecond)
}
