package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptowatt).Yoctowatts())
	assert.Equal(t, 1000.0000000000001, (1 * Attowatt).Zeptowatts()) // round error, expected 1e21
	assert.Equal(t, 1e3, (1 * Femtowatt).Attowatts())
	assert.Equal(t, 999.9999999999999, (1 * Picowatt).Femtowatts()) // round error, expeced 1e15
	assert.Equal(t, 1000.0000000000001, (1 * Nanowatt).Picowatts()) // round error, expeced 1e15
	assert.Equal(t, 999.9999999999999, (1 * Microwatt).Nanowatts()) // round error, expected 1e9
	assert.Equal(t, 1000.0000000000001, (1 * Milliwatt).Microwatts())

	assert.Equal(t, 1e3, (1 * Watt).Milliwatts())
	assert.Equal(t, 1e2, (1 * Watt).Centiwatts())
	assert.Equal(t, 1e1, (1 * Watt).Deciwatts())
	assert.Equal(t, 1e0, (1 * Watt).Watts())
	assert.Equal(t, 1e-1, (1 * Watt).Decawatts())
	assert.Equal(t, 1e-2, (1 * Watt).Hectowatts())
	assert.Equal(t, 1e-3, (1 * Watt).Kilowatts())

	assert.Equal(t, 1e-3, (1 * Kilowatt).Megawatts())
	assert.Equal(t, 1e-3, (1 * Megawatt).Gigawatts())
	assert.Equal(t, 1e-3, (1 * Gigawatt).Terawatts())
	assert.Equal(t, 1e-3, (1 * Terawatt).Petawatts())
	assert.Equal(t, 1e-3, (1 * Petawatt).Exawatts())
	assert.Equal(t, 1e-3, (1 * Exawatt).Zettawatts())
	assert.Equal(t, 1e-3, (1 * Zettawatt).Yottawatts()) // round error, expected 1e-24

	// non-SI
	assert.Equal(t, 13.596216173039045, (10000 * Watt).Pferdestarke())
}
