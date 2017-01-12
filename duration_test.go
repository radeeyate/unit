package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {

	// SI
	assert.Equal(t, 1e3, (1 * Zeptosecond).Yoctoseconds())
	assert.Equal(t, 1000.0000000000001, (1 * Attosecond).Zeptoseconds()) // round error, expected 1e3
	assert.Equal(t, 1e3, (1 * Femtosecond).Attoseconds())
	assert.Equal(t, 999.9999999999999, (1 * Picosecond).Femtoseconds())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Nanosecond).Picoseconds())   // round error, expected 1e3
	assert.Equal(t, 999.9999999999999, (1 * Microsecond).Nanoseconds())   // round error, expected 1e3
	assert.Equal(t, 1000.0000000000001, (1 * Millisecond).Microseconds()) // round error, expected 1e3

	assert.Equal(t, 1e3, (1 * Second).Milliseconds())
	assert.Equal(t, 1e2, (1 * Second).Centiseconds())
	assert.Equal(t, 1e1, (1 * Second).Deciseconds())
	assert.Equal(t, 1e0, (1 * Second).Seconds())
	assert.Equal(t, 1e-1, (1 * Second).Decaseconds())
	assert.Equal(t, 1e-2, (1 * Second).Hectoseconds())
	assert.Equal(t, 1e-3, (1 * Second).Kiloseconds())

	assert.Equal(t, 1e-3, (1 * Kilosecond).Megaseconds())
	assert.Equal(t, 1e-3, (1 * Megasecond).Gigaseconds())
	assert.Equal(t, 1e-3, (1 * Gigasecond).Teraseconds())
	assert.Equal(t, 1e-3, (1 * Terasecond).Petaseconds())
	assert.Equal(t, 1e-3, (1 * Petasecond).Exaseconds())
	assert.Equal(t, 1e-3, (1 * Exasecond).Zettaseconds())
	assert.Equal(t, 1e-3, (1 * Zettasecond).Yottaseconds())

	// non-SI
	assert.Equal(t, 0.016666666666666666, (1 * Second).Minutes())
	assert.Equal(t, 0.0002777777777777778, (1 * Second).Hours())
	assert.Equal(t, 1.1574074074074073e-5, (1 * Second).Days())
	assert.Equal(t, 1.6534391534391535e-6, (1 * Second).Weeks())
	assert.Equal(t, 3.8580246913580245e-7, (1 * Second).ThirtyDayMonths())
	assert.Equal(t, 3.168808781402895e-8, (1 * Second).JulianYears())
}
