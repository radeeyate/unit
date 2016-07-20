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

	// imperial liquid
	assert.Equal(t, 277.41943279162155, (1 * ImperialGallon).CubicInches())
	assert.Equal(t, 1.1365225, (1 * ImperialQuart).Liters())
	assert.Equal(t, 568.2612500000001, (1 * ImperialPint).Milliliters())  // round error, expected 568.26125
	assert.Equal(t, 142.06531250000003, (1 * ImperialGill).Milliliters()) // round error, expected 142.0653125
	assert.Equal(t, 284.13062500000007, (1 * ImperialCup).Milliliters())
	assert.Equal(t, 28.413062500000006, (1 * ImperialFluidOunce).Milliliters()) // round error, expected 28.4130625
	assert.Equal(t, 3.5516328125000007, (1 * ImperialFluidDram).Milliliters())  // round error, expected 3.5516328125
	assert.Equal(t, 9.09218, (1 * ImperialPeck).Liters())
	assert.Equal(t, 36.36872, (1 * ImperialBushel).Liters())

	// US liquid
	assert.Equal(t, 3.7854117839999994, (1 * USLiquidGallon).Liters()) // round error, expected 3.785411784
	assert.Equal(t, 0.9463529459999999, (1 * USLiquidQuart).Liters())  // round error, expected 0.946352946
	assert.Equal(t, 473.176473, (1 * USLiquidPint).Milliliters())
	assert.Equal(t, 236.5882365, (1 * USCup).Milliliters())
	assert.Equal(t, 8.115365448442319, (1 * USLegalCup).USFluidOunces())
	assert.Equal(t, 32.0, (1 * USGill).USFluidDrams())
	assert.Equal(t, 14.78676478125, (1 * USTableSpoon).Milliliters())
	assert.Equal(t, 4.92892159375, (1 * USTeaSpoon).Milliliters())
	assert.Equal(t, 3.6966911953125, (1 * USFluidDram).Milliliters())
	assert.Equal(t, 29.5735295625, (1 * USFluidOunce).Milliliters())

	// US dry
	assert.Equal(t, 1.101220942715, (1 * USDryQuart).Liters())
	assert.Equal(t, 35.23907016688, (1 * USBushel).Liters())
	assert.Equal(t, 8.80976754172, (1 * USPeck).Liters())
	assert.Equal(t, 4.40488377086, (1 * USDryGallon).Liters())
	assert.Equal(t, 550.6104713575, (1 * USDryPint).Milliliters())
}
