package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngle(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptoradian).Yoctoradians())
	assert.Equal(t, 1000.0000000000001, (1 * Attoradian).Zeptoradians()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtoradian).Attoradians())
	assert.Equal(t, 999.9999999999999, (1 * Picoradian).Femtoradians())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanoradian).Picoradians())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microradian).Nanoradians())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Milliradian).Microradians()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Radian).Milliradians())
	assert.Equal(t, 1e2, (1 * Radian).Centiradians())
	assert.Equal(t, 1e1, (1 * Radian).Deciradians())
	assert.Equal(t, 1e0, (1 * Radian).Radians())

	// constant ~57.2958 from https://en.wikipedia.org/wiki/Radian#Conversion_between_radians_and_degrees
	assert.Equal(t, 57.295779513082321, (1 * Radian).Degrees())

	// constants from https://en.wikipedia.org/wiki/Minute_and_second_of_arc#Symbols_and_abbreviations
	assert.Equal(t, 17.453292519943295, (1 * Degree).Milliradians())
	assert.Equal(t, 290.8882086657216, (1 * Arcminute).Microradians())
	assert.Equal(t, 4.84813681109536, (1 * Arcsecond).Microradians())

	assert.Equal(t, 60.0, (1 * Degree).Arcminutes())
	assert.Equal(t, 60.0, (1 * Arcminute).Arcseconds())
	assert.Equal(t, 999.9999999999999, (1 * Arcsecond).Milliarcseconds()) // round error, expected 1e3
	assert.Equal(t, 1e6, (1 * Arcsecond).Microarcseconds())
}
