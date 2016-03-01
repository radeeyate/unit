package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatarate(t *testing.T) {

	// base 10 (SI prefixes)
	assert.Equal(t, 1e0, (1 * BitPerSecond).BitsPerSecond())
	assert.Equal(t, 1e-3, (1 * BitPerSecond).KilobitsPerSecond())
	assert.Equal(t, 1e-3, (1 * KilobitPerSecond).MegabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * MegabitPerSecond).GigabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * GigabitPerSecond).TerabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * TerabitPerSecond).PetabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * PetabitPerSecond).ExabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * ExabitPerSecond).ZettabitsPerSecond())
	assert.Equal(t, 1e-3, (1 * ZettabitPerSecond).YottabitsPerSecond())

	assert.Equal(t, 1e0, (1 * BytePerSecond).BytesPerSecond())
	assert.Equal(t, 1e-3, (1 * BytePerSecond).KilobytesPerSecond())
	assert.Equal(t, 1e-3, (1 * KilobytePerSecond).MegabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * MegabytePerSecond).GigabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * GigabytePerSecond).TerabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * TerabytePerSecond).PetabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * PetabytePerSecond).ExabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * ExabytePerSecond).ZettabytesPerSecond())
	assert.Equal(t, 1e-3, (1 * ZettabytePerSecond).YottabytesPerSecond())

	// base 2 (IEC prefixes)
	assert.Equal(t, 0.0009765625, (1 * BitPerSecond).KibibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * KibibitPerSecond).MebibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * MebibitPerSecond).GibibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * GibibitPerSecond).TebibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * TebibitPerSecond).PebibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * PebibitPerSecond).ExbibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * ExbibitPerSecond).ZebibitsPerSecond())
	assert.Equal(t, 0.0009765625, (1 * ZebibitPerSecond).YobibitsPerSecond())

	assert.Equal(t, 0.0009765625, (1 * BytePerSecond).KibibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * KibibytePerSecond).MebibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * MebibytePerSecond).GibibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * GibibytePerSecond).TebibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * TebibytePerSecond).PebibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * PebibytePerSecond).ExbibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * ExbibytePerSecond).ZebibytesPerSecond())
	assert.Equal(t, 0.0009765625, (1 * ZebibytePerSecond).YobibytesPerSecond())
}
