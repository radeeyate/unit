package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAcceleration(t *testing.T) {

	// constants from https://en.wikipedia.org/wiki/Standard_gravity#Conversions
	assert.Equal(t, 0.03280839895013123, (1 * Gal).FeetPerSecondSquared())
	assert.Equal(t, 0.01, (1 * Gal).MetersPerSecondSquared())
	assert.Equal(t, 1e0, (1 * Gal).CentimetersPerSecondSquared()) // alias
	assert.Equal(t, 0.0010197162129779284, (1 * Gal).StandardGravity())

	assert.Equal(t, 30.4800, (1 * FootPerSecondSquared).Gals())
	assert.Equal(t, 0.304800, (1 * FootPerSecondSquared).MetersPerSecondSquared())
	assert.Equal(t, 0.031080950171567256, (1 * FootPerSecondSquared).StandardGravity())

	assert.Equal(t, 100.0, (1 * MeterPerSecondSquared).Gals())
	assert.Equal(t, 3.280839895013123, (1 * MeterPerSecondSquared).FeetPerSecondSquared())
	assert.Equal(t, 0.10197162129779283, (1 * MeterPerSecondSquared).StandardGravity())

	assert.Equal(t, 980.665, (1 * StandardGravity).Gals())
	assert.Equal(t, 32.17404855643044, (1 * StandardGravity).FeetPerSecondSquared())
	assert.Equal(t, 9.80665, (1 * StandardGravity).MetersPerSecondSquared())
}
