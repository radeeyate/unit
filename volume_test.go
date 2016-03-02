package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolume(t *testing.T) {

	// SI derived
	assert.Equal(t, 1000.0000000000001, (1 * Zepoliter).Yoctoliters()) // round error, expected 1e3
	assert.Equal(t, 1000.0000000000002, (1 * Attoliter).Zepoliters())  // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtoliter).Attoliters())
	assert.Equal(t, 1e3, (1 * Picoliter).Femtoliters())
	assert.Equal(t, 1000.0000000000001, (1 * Nanoliter).Picoliters())  // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microliter).Nanoliters())  // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Milliliter).Microliters()) // round error, expected 1e3

	assert.Equal(t, 1000.0000000000001, (1 * Liter).Milliliters()) // round error, expected 1e3
	assert.Equal(t, 1e2, (1 * Liter).Centiliters())
	assert.Equal(t, 1e1, (1 * Liter).Deciliters())
	assert.Equal(t, 1e0, (1 * Liter).Liters())
	assert.Equal(t, 1e-1, (1 * Liter).Decaliters())
	assert.Equal(t, 1e-2, (1 * Liter).Hectoliters())
	assert.Equal(t, 1e-3, (1 * Liter).Kiloliters())

	assert.Equal(t, 1e-3, (1 * Kiloliter).Megaliters())
	assert.Equal(t, 1e-3, (1 * Megaliter).Gigaliters())
	assert.Equal(t, 1e-3, (1 * Gigaliter).Teraliters())
	assert.Equal(t, 1e-3, (1 * Teraliter).Petaliters())
	assert.Equal(t, 1e-3, (1 * Petaliter).Exaliters())
	assert.Equal(t, 1e-3, (1 * Exaliter).Zettaliters())

	// SI
	assert.Equal(t, 1.0000000000000001e+9, (1 * CubicZeptometer).CubicYoctometers()) // round error, expected 1e9
	assert.Equal(t, 1e9, (1 * CubicAttometer).CubicZeptometers())
	assert.Equal(t, 1e9, (1 * CubicFemtometer).CubicAttometers())
	assert.Equal(t, 1e9, (1 * CubicPicometer).CubicFemtometers())
	assert.Equal(t, 1.0000000000000001e+9, (1 * CubicNanometer).CubicPicometers()) // round error, expected 1e9
	assert.Equal(t, 1e9, (1 * CubicMicrometer).CubicNanometers())
	assert.Equal(t, 1e9, (1 * CubicMillimeter).CubicMicrometers())

	assert.Equal(t, 9.999999999999999e+8, (1 * CubicMeter).CubicMillimeters()) // round error, expected 1e9
	assert.Equal(t, 1e6, (1 * CubicMeter).CubicCentimeters())
	assert.Equal(t, 1e3, (1 * CubicMeter).CubicDecimeters())

	assert.Equal(t, 1e-3, (1 * Liter).CubicMeters())
	assert.Equal(t, 1e0, (1 * CubicMeter).CubicMeters())

	assert.Equal(t, 1e-3, (1 * CubicMeter).CubicDecameters())
	assert.Equal(t, 1e-6, (1 * CubicMeter).CubicHectometers())
	assert.Equal(t, 1e-9, (1 * CubicMeter).CubicKilometers())

	assert.Equal(t, 1e-9, (1 * CubicKilometer).CubicMegameters())
	assert.Equal(t, 1e-9, (1 * CubicMegameter).CubicGigameters())
	assert.Equal(t, 1e-9, (1 * CubicGigameter).CubicTerameters())
	assert.Equal(t, 1e-9, (1 * CubicTerameter).CubicPetameters())
	assert.Equal(t, 9.999999999999999e-10, (1 * CubicPetameter).CubicExameters()) // round error, expected 1e-9
	assert.Equal(t, 1e-9, (1 * CubicExameter).CubicZettameters())
	assert.Equal(t, 1e-9, (1 * CubicZettameter).CubicYottameters())

	// US
	assert.Equal(t, 1728.0, (1 * CubicFoot).CubicInches())
	assert.Equal(t, 1e0, (1 * CubicFoot).CubicFeet())
	assert.Equal(t, 0.037037037037037035, (1 * CubicFoot).CubicYards())
	assert.Equal(t, 1.8342646506386176e-10, (1 * CubicYard).CubicMiles())
	assert.Equal(t, 9.391362885602761e-8, (1 * CubicYard).CubicFurlongs())
}
