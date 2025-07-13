package unit

import "math"

// Value represents a value with a unit
type Value struct {
	val  float64
	unit theUnit
	base float64
}

// NewConverter creates a new Value for conversion
func NewConverter(v float64) Value {
	return Value{val: v}
}

// From sets the source unit for the conversion
func (v Value) From(symbol string) Value {
	multiplier, unitSymbol := parseSymbol(symbol)
	if multiplier == 0 {
		// Fallback to non-SI units
		result, ok := fromMap[symbol]
		if !ok {
			return Value{}
		}
		return result(v.val)
	}

	unit, ok := siUnits[unitSymbol]
	if !ok {
		return Value{}
	}

	return Value{val: v.val, unit: unit.unit, base: multiplier}
}

// To performs the conversion to the target unit
func (v Value) To(symbol string) (float64, error) {
	if v.unit == 0 {
		return 0, ErrNotFound
	}

	multiplier, unitSymbol := parseSymbol(symbol)
	if multiplier == 0 {
		// Fallback to non-SI units
		result, ok := toMap[symbol]
		if !ok {
			return 0, ErrNotFound
		}
		return result(v)
	}

	unit, ok := siUnits[unitSymbol]
	if !ok || unit.unit != v.unit {
		return 0, ErrConversion
	}

	result := v.val * v.base / multiplier
	return round(result, 12), nil
}

func round(val float64, precision int) float64 {
	if precision < 0 {
		precision = 0
	}
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
