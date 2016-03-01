package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnergy(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptojoule).Yoctojoules())
	assert.Equal(t, 1000.0000000000001, (1 * Attojoule).Zeptojoules()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtojoule).Attojoules())
	assert.Equal(t, 999.9999999999999, (1 * Picojoule).Femtojoules())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanojoule).Picojoules())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microjoule).Nanojoules())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millijoule).Microjoules()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Joule).Millijoules())
	assert.Equal(t, 1e2, (1 * Joule).Centijoules())
	assert.Equal(t, 1e1, (1 * Joule).Decijoules())
	assert.Equal(t, 1e0, (1 * Joule).Joules())
	assert.Equal(t, 1e-1, (1 * Joule).Decajoules())
	assert.Equal(t, 1e-2, (1 * Joule).Hectojoules())
	assert.Equal(t, 1e-3, (1 * Joule).Kilojoules())

	assert.Equal(t, 1e-3, (1 * Kilojoule).Megajoules())
	assert.Equal(t, 1e-3, (1 * Megajoule).Gigajoules())
	assert.Equal(t, 1e-3, (1 * Gigajoule).Terajoules())
	assert.Equal(t, 1e-3, (1 * Terajoule).Petajoules())
	assert.Equal(t, 1e-3, (1 * Petajoule).Exajoules())
	assert.Equal(t, 1e-3, (1 * Exajoule).Zettajoules())
	assert.Equal(t, 1e-3, (1 * Zettajoule).Yottajoules())

	// watt hours
	assert.Equal(t, 999.9999999999999, (1 * ZeptowattHour).YoctowattHours()) // round error, expected 1e3
	assert.Equal(t, 1000.0000000000002, (1 * AttowattHour).ZeptowattHours()) // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * FemtowattHour).AttowattHours()) // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * PicowattHour).FemtowattHours())  // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * NanowattHour).PicowattHours())  // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * MicrowattHour).NanowattHours())  // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * MilliwattHour).MicrowattHours())

	assert.Equal(t, 1e3, (1 * WattHour).MilliwattHours())
	assert.Equal(t, 1e2, (1 * WattHour).CentiwattHours())
	assert.Equal(t, 1e1, (1 * WattHour).DeciwattHours())

	assert.Equal(t, 3600.0, (1 * WattHour).Joules())
	assert.Equal(t, 1e0, (1 * WattHour).WattHours())

	assert.Equal(t, 1e-1, (1 * WattHour).DecawattHours())
	assert.Equal(t, 1e-2, (1 * WattHour).HectowattHours())
	assert.Equal(t, 1e-3, (1 * WattHour).KilowattHours())

	assert.Equal(t, 1e-3, (1 * KilowattHour).MegawattHours())
	assert.Equal(t, 1e-3, (1 * MegawattHour).GigawattHours())
	assert.Equal(t, 1e-3, (1 * GigawattHour).TerawattHours())
	assert.Equal(t, 1e-3, (1 * TerawattHour).PetawattHours())
	assert.Equal(t, 1e-3, (1 * PetawattHour).ExawattHours())
	assert.Equal(t, 1e-3, (1 * ExawattHour).ZettawattHours())
	assert.Equal(t, 1e-3, (1 * ZettawattHour).YottawattHours())

	// calories
	assert.Equal(t, 4.184, (1 * Gramcalorie).Joules())
	assert.Equal(t, 4184.0, (1 * Kilocalorie).Joules())
	assert.Equal(t, 4.184e+6, (1 * Megacalorie).Joules())
}
