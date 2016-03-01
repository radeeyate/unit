package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForce(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Force#Units_of_measurement
	assert.Equal(t, 99999.99999999999, (1 * Newton).Dynes()) // round error, expected 1e5
	assert.Equal(t, 0.10197162129779283, (1 * Newton).KilogramForce())
	assert.Equal(t, 0.22480892365533914, (1 * Newton).PoundForce())
	assert.Equal(t, 7.233011464323171, (1 * Newton).Poundals())

	assert.Equal(t, 1e-5, (1 * Dyne).Newtons())
	assert.Equal(t, 1.0197162129779284e-6, (1 * Dyne).KilogramForce())
	assert.Equal(t, 2.2480892365533917e-6, (1 * Dyne).PoundForce())
	assert.Equal(t, 7.233011464323172e-05, (1 * Dyne).Poundals())

	assert.Equal(t, 9.80665, (1 * KilogramForce).Newtons())
	assert.Equal(t, 980664.9999999999, (1 * KilogramForce).Dynes()) // round error, expected 980665
	assert.Equal(t, 2.204622431164631, (1 * KilogramForce).PoundForce())
	assert.Equal(t, 70.93161187660482, (1 * KilogramForce).Poundals())

	assert.Equal(t, 4.448222, (1 * PoundForce).Newtons())
	assert.Equal(t, 444822.2, (1 * PoundForce).Dynes())
	assert.Equal(t, 0.45359240923251065, (1 * PoundForce).KilogramForce())
	assert.Equal(t, 32.174040721854546, (1 * PoundForce).Poundals())

	assert.Equal(t, 0.138255, (1 * Poundal).Newtons())
	assert.Equal(t, 13825.499999999998, (1 * Poundal).Dynes()) // round error, expected 13825.5
	assert.Equal(t, 0.014098086502526346, (1 * Poundal).KilogramForce())
	assert.Equal(t, 0.03108095773996891, (1 * Poundal).PoundForce())
}
