package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolume(t *testing.T) {

	// SI derived
	assert.Equal(t, 1000.0000000000001, (1 * Zepoliter).Yoctoliters())
	assert.Equal(t, 1000.0000000000002, (1 * Attoliter).Zepoliters())
	assert.Equal(t, 1e3, (1 * Femtoliter).Attoliters())
	assert.Equal(t, 1e3, (1 * Picoliter).Femtoliters())
	assert.Equal(t, 1000.0000000000001, (1 * Nanoliter).Picoliters())
	assert.Equal(t, 999.9999999999999, (1 * Microliter).Nanoliters())
	assert.Equal(t, 999.9999999999999, (1 * Milliliter).Microliters())

	assert.Equal(t, 1000.0000000000001, (1 * Liter).Milliliters())
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
	assert.Equal(t, 1.0000000000000001e+9, (1 * CubicZeptometer).CubicYoctometers())
	assert.Equal(t, 1e9, (1 * CubicAttometer).CubicZeptometers())
	assert.Equal(t, 1e9, (1 * CubicFemtometer).CubicAttometers())
	assert.Equal(t, 1e9, (1 * CubicPicometer).CubicFemtometers())
	assert.Equal(t, 1.0000000000000001e+9, (1 * CubicNanometer).CubicPicometers())
	assert.Equal(t, 1e9, (1 * CubicMicrometer).CubicNanometers())
	assert.Equal(t, 1e9, (1 * CubicMillimeter).CubicMicrometers())

	assert.Equal(t, 9.999999999999999e+8, (1 * CubicMeter).CubicMillimeters())
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
	assert.Equal(t, 9.999999999999999e-10, (1 * CubicPetameter).CubicExameters())
	assert.Equal(t, 1e-9, (1 * CubicExameter).CubicZettameters())
	assert.Equal(t, 1e-9, (1 * CubicZettameter).CubicYottameters())

	// US
	assert.Equal(t, 1728.0, (1 * CubicFoot).CubicInches())
	assert.Equal(t, 1e0, (1 * CubicFoot).CubicFeet())
	assert.Equal(t, 0.037037037037037035, (1 * CubicFoot).CubicYards())
	assert.Equal(t, 1.8342646506386176e-10, (1 * CubicYard).CubicMiles())
	assert.Equal(t, 9.391362885602761e-8, (1 * CubicYard).CubicFurlongs())

	// imperial liquid
	assert.Equal(t, 0.21996924829908776, (1 * Liter).ImperialGallons())
	assert.Equal(t, 0.879876993196351, (1 * Liter).ImperialQuarts())
	assert.Equal(t, 1.759753986392702, (1 * Liter).ImperialPints())
	assert.Equal(t, 7.039015945570808, (1 * Liter).ImperialGills())
	assert.Equal(t, 3.519507972785404, (1 * Liter).ImperialCups())
	assert.Equal(t, 35.19507972785404, (1 * Liter).ImperialFluidOunces())
	assert.Equal(t, 281.56063782283235, (1 * Liter).ImperialFluidDrams())
	assert.Equal(t, 0.10998462414954388, (1 * Liter).ImperialPecks())
	assert.Equal(t, 0.02749615603738597, (1 * Liter).ImperialBushels())

	// US liquid
	assert.Equal(t, 0.26417205235814845, (1 * Liter).USLiquidGallons())
	assert.Equal(t, 1.0566882094325938, (1 * Liter).USLiquidQuarts())
	assert.Equal(t, 2.1133764188651876, (1 * Liter).USLiquidPints())
	assert.Equal(t, 4.226752837730375, (1 * Liter).USCups())
	assert.Equal(t, 4.166666666666667, (1 * Liter).USLegalCups())
	assert.Equal(t, 8.45350567546075, (1 * Liter).USGills())
	assert.Equal(t, 67.628045403686, (1 * Liter).USTableSpoons())
	assert.Equal(t, 202.884136211058, (1 * Liter).USTeaSpoons())
	assert.Equal(t, 270.512181614744, (1 * Liter).USFluidDrams())
	assert.Equal(t, 33.814022701843, (1 * Liter).USFluidOunces())

	// US dry
	assert.Equal(t, 0.9080829842688559, (1 * Liter).USDryQuarts())
	assert.Equal(t, 0.028377593258401747, (1 * Liter).USBushels())
	assert.Equal(t, 0.11351037303360699, (1 * Liter).USPecks())
	assert.Equal(t, 0.22702074606721398, (1 * Liter).USDryGallons())
	assert.Equal(t, 1.8161659685377118, (1 * Liter).USDryPints())
}
