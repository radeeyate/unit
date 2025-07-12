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
	assertFloatEqual(t, 1e3, (1 * Millifarad).Microfarads())

	assertFloatEqual(t, 1e3, (1 * Farad).Millifarads())
	assertFloatEqual(t, 1e2, (1 * Farad).Decifarads())
	assertFloatEqual(t, 1e1, (1 * Farad).Centifarads())
	assertFloatEqual(t, 1e0, (1 * Farad).Farads())
	assertFloatEqual(t, 1e-1, (1 * Farad).Decafarads())
	assertFloatEqual(t, 1e-2, (1 * Farad).Hectofarads())
	assertFloatEqual(t, 1e-3, (1 * Farad).Kilofarads())

	assertFloatEqual(t, 1e-3, (1 * Kilofarad).Megafarads())
	assertFloatEqual(t, 1e-3, (1 * Megafarad).Gigafarads())
	assertFloatEqual(t, 1e-3, (1 * Gigafarad).Terafarads())
	assertFloatEqual(t, 1e-3, (1 * Terafarad).Petafarads())
	assertFloatEqual(t, 1e-3, (1 * Petafarad).Exafarads())
	assertFloatEqual(t, 1e-3, (1 * Exafarad).Zettafarads())
	assertFloatEqual(t, 1e-3, (1 * Zettafarad).Yottafarads())
}
