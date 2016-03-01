package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeed(t *testing.T) {

	// SI
	assert.Equal(t, 10.0, (10 * MetersPerSecond).MetersPerSecond())
	assert.Equal(t, 35.99997120002304, (10 * MetersPerSecond).KilometersPerHour())

	// US
	assert.Equal(t, 32.808398950131235, (10 * MetersPerSecond).FeetPerSecond())
	assert.Equal(t, 22.369362920544024, (10 * MetersPerSecond).MilesPerHour())

	// misc
	assert.Equal(t, 19.438461717893492, (10 * MetersPerSecond).Knots())

	// space
	assert.Equal(t, 0.03335640951981521, (10000000 * MetersPerSecond).SpeedOfLight())
}
