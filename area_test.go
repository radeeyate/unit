package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {

	// SI
	assert.Equal(t, 1.0000000000000001e6, (1 * SquareZeptometer).SquareYoctometers()) // round error, expected 1e6
	assert.Equal(t, 999999.9999999999, (1 * SquareAttometer).SquareZeptometers())     // round error, expected 1e6
	assert.Equal(t, 1.0000000000000001e6, (1 * SquareFemtometer).SquareAttometers())  // round error, expected 1e6
	assert.Equal(t, 999999.9999999999, (1 * SquarePicometer).SquareFemtometers())     // round error, expected 1e6
	assert.Equal(t, 1.0000000000000001e6, (1 * SquareNanometer).SquarePicometers())   // round error, expected 1e6
	assert.Equal(t, 999999.9999999999, (1 * SquareMicrometer).SquareNanometers())     // round error, expected 1e6
	assert.Equal(t, 1e6, (1 * SquareMillimeter).SquareMicrometers())
	assert.Equal(t, 1e6, (1 * SquareMeter).SquareMillimeters())

	assert.Equal(t, 1e4, (1 * SquareMeter).SquareCentimeters())
	assert.Equal(t, 1e2, (1 * SquareMeter).SquareDecimeters())
	assert.Equal(t, 1e0, (1 * SquareMeter).SquareMeters())
	assert.Equal(t, 1e-2, (1 * SquareMeter).SquareDecameter())
	assert.Equal(t, 1e-4, (1 * SquareMeter).SquareHectometer())

	assert.Equal(t, 1e-6, (1 * SquareMeter).SquareKilometers())
	assert.Equal(t, 1e-6, (1 * SquareKilometer).SquareMegameters())
	assert.Equal(t, 1e-6, (1 * SquareMegameter).SquareGigameters())
	assert.Equal(t, 1e-6, (1 * SquareGigameter).SquareTerameters())
	assert.Equal(t, 1e-6, (1 * SquareTerameter).SquarePetameters())
	assert.Equal(t, 1e-6, (1 * SquarePetameter).SquareExameters())
	assert.Equal(t, 1e-6, (1 * SquareExameter).SquareZettameters())
	assert.Equal(t, 1e-6, (1 * SquareZettameter).SquareYottameters())

	// US
	assert.Equal(t, 10.763910416709724, (1 * SquareMeter).SquareFeet())
	assert.Equal(t, 144.0, (1 * SquareFoot).SquareInches())
	assert.Equal(t, 9.0, (1 * SquareYard).SquareFeet())
	assert.Equal(t, 4840.0, (1 * Acre).SquareYards())
	assert.Equal(t, 640.0, (1 * SquareMile).Acres())
	assert.Equal(t, 0.0015624999999999999, (1 * Acre).SquareMiles())

	// aliases
	assert.Equal(t, 1.0, (1 * SquareMeter).Centiares())
	assert.Equal(t, 1.0, (1 * SquareDecameter).Ares())
	assert.Equal(t, 1.0, (1 * SquareHectometer).Hectares())
}
