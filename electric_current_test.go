package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElectricCurrent(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptoampere).Yoctoamperes())
	assert.Equal(t, 1000.0000000000001, (1 * Attoampere).Zeptoamperes()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtoampere).Attoamperes())
	assert.Equal(t, 999.9999999999999, (1 * Picoampere).Femtoamperes())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanoampere).Picoamperes())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microampere).Nanoamperes())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Milliampere).Microamperes()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Ampere).Milliamperes())
	assert.Equal(t, 1e2, (1 * Ampere).Deciamperes())
	assert.Equal(t, 1e1, (1 * Ampere).Centiamperes())
	assert.Equal(t, 1e0, (1 * Ampere).Amperes())
	assert.Equal(t, 1e-1, (1 * Ampere).Decaamperes())
	assert.Equal(t, 1e-2, (1 * Ampere).Hectoamperes())
	assert.Equal(t, 1e-3, (1 * Ampere).Kiloamperes())

	assert.Equal(t, 1e-3, (1 * Kiloampere).Megaamperes())
	assert.Equal(t, 1e-3, (1 * Megaampere).Gigaamperes())
	assert.Equal(t, 1e-3, (1 * Gigaampere).Teraamperes())
	assert.Equal(t, 1e-3, (1 * Teraampere).Petaamperes())
	assert.Equal(t, 1e-3, (1 * Petaampere).Exaamperes())
	assert.Equal(t, 1e-3, (1 * Exaampere).Zettaamperes())
	assert.Equal(t, 1e-3, (1 * Zettaampere).Yottaamperes())
}
