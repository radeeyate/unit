package unit

import (
	"testing"
)

func TestElectricalInductance(t *testing.T) {

	// SI
	assertFloatEqual(t, 1e3, (1 * Zeptohenry).Yoctohenries())
	assertFloatEqual(t, 1e3, (1 * Attohenry).Zeptohenries())
	assertFloatEqual(t, 1e3, (1 * Femtohenry).Attohenries())
	assertFloatEqual(t, 1e3, (1 * Picohenry).Femtohenries())
	assertFloatEqual(t, 1e3, (1 * Nanohenry).Picohenries())
	assertFloatEqual(t, 1e3, (1 * Microhenry).Nanohenries())
	assertFloatEqual(t, 1e3, (1 * Millihenry).Microhenries())

	assertFloatEqual(t, 1e3, (1 * Henry).Millihenries())
	assertFloatEqual(t, 1e2, (1 * Henry).Decihenries())
	assertFloatEqual(t, 1e1, (1 * Henry).Centihenries())
	assertFloatEqual(t, 1e0, (1 * Henry).Henries())
	assertFloatEqual(t, 1e-1, (1 * Henry).Decahenries())
	assertFloatEqual(t, 1e-2, (1 * Henry).Hectohenries())
	assertFloatEqual(t, 1e-3, (1 * Henry).Kilohenries())

	assertFloatEqual(t, 1e-3, (1 * Kilohenry).Megahenries())
	assertFloatEqual(t, 1e-3, (1 * Megahenry).Gigahenries())
	assertFloatEqual(t, 1e-3, (1 * Gigahenry).Terahenries())
	assertFloatEqual(t, 1e-3, (1 * Terahenry).Petahenries())
	assertFloatEqual(t, 1e-3, (1 * Petahenry).Exahenries())
	assertFloatEqual(t, 1e-3, (1 * Exahenry).Zettahenries())
	assertFloatEqual(t, 1e-3, (1 * Zettahenry).Yottahenries())
}
