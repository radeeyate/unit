package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {

	assert.Equal(t, 373.15, FromCelsius(100).Kelvin())
	assert.Equal(t, 310.9277777777778, FromFahrenheit(100).Kelvin())
	assert.Equal(t, 100.0, FromKelvin(100).Kelvin())
	assert.Equal(t, 576.180303030303, FromNewton(100).Kelvin())
	assert.Equal(t, 306.4833333333333, FromDelisle(100).Kelvin())
	assert.Equal(t, 398.15, FromReaumur(100).Kelvin())
	assert.Equal(t, 449.3404761904762, FromRomer(100).Kelvin())
	assert.Equal(t, 55.555555555555515, FromRankine(100).Kelvin())

	assert.Equal(t, -173.14999999999998, FromKelvin(100).Celsius())
	assert.Equal(t, -279.67, FromKelvin(100).Fahrenheit())
	assert.Equal(t, -57.13949999999999, FromKelvin(100).Newton())
	assert.Equal(t, 409.72499999999997, FromKelvin(100).Delisle())
	assert.Equal(t, -83.40374999999999, FromKelvin(100).Romer())
	assert.Equal(t, -138.51999999999998, FromKelvin(100).Reaumur())
	assert.Equal(t, 180.00000000000006, FromKelvin(100).Rankine())
}
