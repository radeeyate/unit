package unit

import (
	"testing"
)

func TestElectricalCapacitance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptofarad).Yoctofarads())
	assertFloatEqual(t, 1e3, (1 * Attofarad).Zeptofarads())
	assertFloatEqual(t, 1e3, (1 * Femtofarad).Attofarads())
	assertFloatEqual(t, 1e3, (1 * Picofarad).Femtofarads())
	assertFloatEqual(t, 1e3, (1 * Nanofarad).Picofarads())
	assertFloatEqual(t, 1e3, (1 * Microfarad).Nanofarads())

	val, err := NewConverter(1).From("F").To("mF")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1000, val)

	val, err = NewConverter(1).From("F").To("kF")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 0.001, val)

	val, err = NewConverter(1).From("mF").To("uF")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1000, val)

	val, err = NewConverter(1).From("F").To("daF")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1e-1, val)

	val, err = NewConverter(1).From("F").To("hF")
	if err != nil {
		t.Errorf("conversion error: %v", err)
	}
	assertFloatEqual(t, 1e-2, val)

	assertFloatEqual(t, 1e-3, (1 * Kilofarad).Megafarads())
	assertFloatEqual(t, 1e-3, (1 * Megafarad).Gigafarads())
	assertFloatEqual(t, 1e-3, (1 * Gigafarad).Terafarads())
	assertFloatEqual(t, 1e-3, (1 * Terafarad).Petafarads())
	assertFloatEqual(t, 1e-3, (1 * Petafarad).Exafarads())
	assertFloatEqual(t, 1e-3, (1 * Exafarad).Zettafarads())
	assertFloatEqual(t, 1e-3, (1 * Zettafarad).Yottafarads())
}
