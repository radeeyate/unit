package unit

import (
	"testing"
)

func TestElectricalResistance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e0, (1 * Ohm).Ohms())

	// Test conversions
	val, err := NewConverter(1).From("kΩ").To("Ω")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1000, val)

	val, err = NewConverter(1).From("Ω").To("mΩ")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1000, val)

	val, err = NewConverter(1).From("µΩ").To("nΩ")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1000, val)
}
