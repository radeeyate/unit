package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequency(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptohertz).Yoctohertz())
	assert.Equal(t, 1000.0000000000001, (1 * Attohertz).Zeptohertz()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtohertz).Attohertz())
	assert.Equal(t, 999.9999999999999, (1 * Picohertz).Femtohertz()) // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanohertz).Picohertz()) // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microhertz).Nanohertz()) // round error, expected 1e3
	assert.Equal(t, 1e6, (1 * Hertz).Microhertz())

	assert.Equal(t, 1e3, (1 * Hertz).Millihertz())
	assert.Equal(t, 1e2, (1 * Hertz).Centihertz())
	assert.Equal(t, 1e1, (1 * Hertz).Decihertz())
	assert.Equal(t, 1e0, (1 * Hertz).Hertz())
	assert.Equal(t, 1e-1, (1 * Hertz).Decahertz())
	assert.Equal(t, 1e-2, (1 * Hertz).Hectohertz())
	assert.Equal(t, 1e-3, (1 * Hertz).Kilohertz())

	assert.Equal(t, 1e-6, (1 * Hertz).Megahertz())
	assert.Equal(t, 1e-3, (1 * Megahertz).Gigahertz())
	assert.Equal(t, 1e-3, (1 * Gigahertz).Terahertz())
	assert.Equal(t, 1e-3, (1 * Terahertz).Petahertz())
	assert.Equal(t, 1e-3, (1 * Petahertz).Exahertz())
	assert.Equal(t, 1e-3, (1 * Exahertz).Zettahertz())
	assert.Equal(t, 1e-3, (1 * Zettahertz).Yottahertz())
}
