package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoltage(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptovolt).Yoctovolts())
	assert.Equal(t, 1000.0000000000001, (1 * Attovolt).Zeptovolts()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtovolt).Attovolts())
	assert.Equal(t, 999.9999999999999, (1 * Picovolt).Femtovolts())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanovolt).Picovolts())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microvolt).Nanovolts())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millivolt).Microvolts()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Volt).Millivolts())
	assert.Equal(t, 1e2, (1 * Volt).Centivolts())
	assert.Equal(t, 1e1, (1 * Volt).Decivolts())
	assert.Equal(t, 1e0, (1 * Volt).Volts())
	assert.Equal(t, 1e-1, (1 * Volt).Decavolts())
	assert.Equal(t, 1e-2, (1 * Volt).Hectovolts())
	assert.Equal(t, 1e-3, (1 * Volt).Kilovolts())

	assert.Equal(t, 1e-3, (1 * Kilovolt).Megavolts())
	assert.Equal(t, 1e-3, (1 * Megavolt).Gigavolts())
	assert.Equal(t, 1e-3, (1 * Gigavolt).Teravolts())
	assert.Equal(t, 1e-3, (1 * Teravolt).Petavolts())
	assert.Equal(t, 1e-3, (1 * Petavolt).Exavolts())
	assert.Equal(t, 1e-3, (1 * Exavolt).Zettavolts())
	assert.Equal(t, 1e-3, (1 * Zettavolt).Yottavolts())
}
