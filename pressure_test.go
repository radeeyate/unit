package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPressure(t *testing.T) {

	// SI derived
	assert.Equal(t, 1e3, (1 * Zeptopascal).Yoctopascals())
	assert.Equal(t, 1000.0000000000001, (1 * Attopascal).Zeptopascals()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtopascal).Attopascals())
	assert.Equal(t, 999.9999999999999, (1 * Picopascal).Femtopascals())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanopascal).Picopascals())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Micropascal).Nanopascals())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millipascal).Micropascals()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Pascal).Millipascals())
	assert.Equal(t, 1e2, (1 * Pascal).Centipascals())
	assert.Equal(t, 1e1, (1 * Pascal).Decipascals())
	assert.Equal(t, 1e0, (1 * Pascal).Pascals())
	assert.Equal(t, 1e-1, (1 * Pascal).Decapascals())
	assert.Equal(t, 1e-2, (1 * Pascal).Hectopascals())
	assert.Equal(t, 1e-3, (1 * Pascal).Kilopascals())

	assert.Equal(t, 1e-3, (1 * Kilopascal).Megapascals())
	assert.Equal(t, 1e-3, (1 * Megapascal).Gigapascals())
	assert.Equal(t, 1e-3, (1 * Gigapascal).Terapascals())
	assert.Equal(t, 1e-3, (1 * Terapascal).Petapascals())
	assert.Equal(t, 1e-3, (1 * Petapascal).Exapascals())
	assert.Equal(t, 1e-3, (1 * Exapascal).Zettapascals())
	assert.Equal(t, 1e-3, (1 * Zettapascal).Yottapascals())

	// non-SI
	assert.Equal(t, 999.9999999999999, (1 * Zeptobar).Yoctobars()) // round error, expected 1e3
	assert.Equal(t, 1000.0000000000002, (1 * Attobar).Zeptobars()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtobar).Attobars())
	assert.Equal(t, 999.9999999999999, (1 * Picobar).Femtobars())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanobar).Picobars())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microbar).Nanobars())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millibar).Microbars()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Bar).Millibars())
	assert.Equal(t, 1e2, (1 * Bar).Centibars())
	assert.Equal(t, 1e1, (1 * Bar).Decibars())
	assert.Equal(t, 1e0, (1 * Bar).Bars())
	assert.Equal(t, 1e-1, (1 * Bar).Decabars())
	assert.Equal(t, 1e-2, (1 * Bar).Hectobars())
	assert.Equal(t, 1e-3, (1 * Bar).Kilobars())

	assert.Equal(t, 1e-3, (1 * Kilobar).Megabars())
	assert.Equal(t, 1e-3, (1 * Megabar).Gigabars())
	assert.Equal(t, 1e-3, (1 * Gigabar).Terabars())
	assert.Equal(t, 1e-3, (1 * Terabar).Petabars())
	assert.Equal(t, 1e-3, (1 * Petabar).Exabars())
	assert.Equal(t, 0.0009999999999999998, (1 * Exabar).Zettabars())   // round error, expected 1e-3
	assert.Equal(t, 0.0010000000000000002, (1 * Zettabar).Yottabars()) // round error, expected 1e-3

	// misc
	assert.Equal(t, 0.09869232667160129, (10000 * Pascal).Atmospheres())
	assert.Equal(t, 0.10197162129779283, (10000 * Pascal).TechAtmospheres())
	assert.Equal(t, 75.00615050434136, (10000 * Pascal).Torrs())
	assert.Equal(t, 1.4503683935719673, (10000 * Pascal).PoundsPerSquareInch())
}
