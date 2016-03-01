package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatasize(t *testing.T) {

	// base 10 (SI prefixes)
	assert.Equal(t, 1e0, (1 * Bit).Bits())
	assert.Equal(t, 1e-3, (1 * Bit).Kilobits())
	assert.Equal(t, 1e-3, (1 * Kilobit).Megabits())
	assert.Equal(t, 1e-3, (1 * Megabit).Gigabits())
	assert.Equal(t, 1e-3, (1 * Gigabit).Terabits())
	assert.Equal(t, 1e-3, (1 * Terabit).Petabits())
	assert.Equal(t, 1e-3, (1 * Petabit).Exabits())
	assert.Equal(t, 1e-3, (1 * Exabit).Zettabits())
	assert.Equal(t, 1e-3, (1 * Zettabit).Yottabits())

	assert.Equal(t, 0.125, (1 * Bit).Bytes())
	assert.Equal(t, 1e0, (1 * Byte).Bytes())

	assert.Equal(t, 1e-3, (1 * Byte).Kilobytes())
	assert.Equal(t, 1e-3, (1 * Kilobyte).Megabytes())
	assert.Equal(t, 1e-3, (1 * Megabyte).Gigabytes())
	assert.Equal(t, 1e-3, (1 * Gigabyte).Terabytes())
	assert.Equal(t, 1e-3, (1 * Terabyte).Petabytes())
	assert.Equal(t, 1e-3, (1 * Petabyte).Exabytes())
	assert.Equal(t, 1e-3, (1 * Exabyte).Zettabytes())
	assert.Equal(t, 1e-3, (1 * Zettabyte).Yottabytes())

	// base 2 (IEC prefixes)
	assert.Equal(t, 0.0009765625, (1 * Bit).Kibibits())
	assert.Equal(t, 0.0009765625, (1 * Kibibit).Mebibits())
	assert.Equal(t, 0.0009765625, (1 * Mebibit).Gibibits())
	assert.Equal(t, 0.0009765625, (1 * Gibibit).Tebibits())
	assert.Equal(t, 0.0009765625, (1 * Tebibit).Pebibits())
	assert.Equal(t, 0.0009765625, (1 * Pebibit).Exbibits())
	assert.Equal(t, 0.0009765625, (1 * Exbibit).Zebibits())
	assert.Equal(t, 0.0009765625, (1 * Zebibit).Yobibits())

	assert.Equal(t, 0.0009765625, (1 * Byte).Kibibytes())
	assert.Equal(t, 0.0009765625, (1 * Kibibyte).Mebibytes())
	assert.Equal(t, 0.0009765625, (1 * Mebibyte).Gibibytes())
	assert.Equal(t, 0.0009765625, (1 * Gibibyte).Tebibytes())
	assert.Equal(t, 0.0009765625, (1 * Tebibyte).Pebibytes())
	assert.Equal(t, 0.0009765625, (1 * Pebibyte).Exbibytes())
	assert.Equal(t, 0.0009765625, (1 * Exbibyte).Zebibytes())
	assert.Equal(t, 0.0009765625, (1 * Zebibyte).Yobibytes())
}
