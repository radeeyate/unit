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

// FromCubicMetersPerSecond return the value to be converted
func FromCubicMetersPerSecond(val float64) Value {
	return Value{val * float64(CubicMeterPerSecond), flow}
}

// toCubicMetersPerSecond return the converted value
func toCubicMetersPerSecond(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).CubicMetersPerSecond(), nil
}

// CubicMetersPerSecond returns the flow rate in m3/s
func (f Flow) CubicMetersPerSecond() float64 {
	return float64(f)
}

// FromCubicFeetPerSecond return the value to be converted
func FromCubicFeetPerSecond(val float64) Value {
	return Value{val * float64(CubicFootPerSecond), flow}
}

// toCubicFeetPerSecond return the converted value
func toCubicFeetPerSecond(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).CubicFeetPerSecond(), nil
}

// CubicFeetPerSecond returns the flow rate in ft3/s
func (f Flow) CubicFeetPerSecond() float64 {
	return float64(f / CubicFootPerSecond)
}

// FromLitersPerSecond return the value to be converted
func FromLitersPerSecond(val float64) Value {
	return Value{val * float64(LiterPerSecond), flow}
}

// toLitersPerSecond return the converted value
func toLitersPerSecond(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).LitersPerSecond(), nil
}

// LitersPerSecond returns the flow rate in l/s
func (f Flow) LitersPerSecond() float64 {
	return float64(f / LiterPerSecond)
}

// FromLitersPerHour return the value to be converted
func FromLitersPerHour(val float64) Value {
	return Value{val * float64(LiterPerHour), flow}
}

// toLitersPerHour return the converted value
func toLitersPerHour(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).LitersPerHour(), nil
}

// LitersPerHour returns the flow rate in l/h
func (f Flow) LitersPerHour() float64 {
	return float64(f / LiterPerHour)
}

// FromLitersPerMinute return the value to be converted
func FromLitersPerMinute(val float64) Value {
	return Value{val * float64(LiterPerMinute), flow}
}

// toLitersPerMinute return the converted value
func toLitersPerMinute(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).LitersPerMinute(), nil
}

// LitersPerMinute returns the flow rate in l/min
func (f Flow) LitersPerMinute() float64 {
	return float64(f / LiterPerMinute)
}

// FromUSLiquidGallonsPerHour return the value to be converted
func FromUSLiquidGallonsPerHour(val float64) Value {
	return Value{val * float64(USLiquidGallonPerHour), flow}
}

// toUSLiquidGallonsPerHour return the converted value
func toUSLiquidGallonsPerHour(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).USLiquidGallonsPerHour(), nil
}

// USLiquidGallonsPerHour returns the flow rate in gal/h
func (f Flow) USLiquidGallonsPerHour() float64 {
	return float64(f / USLiquidGallonPerHour)
}

// FromUSLiquidGallonsPerMinute return the value to be converted
func FromUSLiquidGallonsPerMinute(val float64) Value {
	return Value{val * float64(USLiquidGallonPerMinute), flow}
}

// toUSLiquidGallonsPerMinute return the converted value
func toUSLiquidGallonsPerMinute(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).USLiquidGallonsPerMinute(), nil
}

// USLiquidGallonsPerMinute returns the flow rate in gal/min
func (f Flow) USLiquidGallonsPerMinute() float64 {
	return float64(f / USLiquidGallonPerMinute)
}

// FromUSLiquidGallonsPerSecond return the value to be converted
func FromUSLiquidGallonsPerSecond(val float64) Value {
	return Value{val * float64(USLiquidGallonPerSecond), flow}
}

// toUSLiquidGallonsPerSecond return the converted value
func toUSLiquidGallonsPerSecond(value Value) (float64, error) {
	if value.unit != flow {
		return 0, ErrConversion
	}
	return Flow(value.val).USLiquidGallonsPerSecond(), nil
}

// USLiquidGallonsPerSecond returns the flow rate in gal/s
func (f Flow) USLiquidGallonsPerSecond() float64 {
	return float64(f / USLiquidGallonPerSecond)
}
