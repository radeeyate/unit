package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMass(t *testing.T) {

	// SI
	assert.Equal(t, 1e3+0.0000000000001, (1 * Zeptogram).Yoctograms())
	assert.Equal(t, 1e3+0.0000000000002, (1 * Attogram).Zeptograms())
	assert.Equal(t, 1e3, (1 * Femtogram).Attograms())
	assert.Equal(t, 1e3, (1 * Picogram).Femtograms())
	assert.Equal(t, 1e3+0.0000000000001, (1 * Nanogram).Picograms())
	assert.Equal(t, 1e3-0.0000000000001, (1 * Microgram).Nanograms())
	assert.Equal(t, 1e3-0.0000000000001, (1 * Milligram).Micrograms())

	assert.Equal(t, 1e3+0.0000000000001, (1 * Gram).Milligrams())
	assert.Equal(t, 1e2, (1 * Gram).Centigrams())
	assert.Equal(t, 1e1, (1 * Gram).Decigrams())
	assert.Equal(t, 1e0, (1 * Gram).Grams())
	assert.Equal(t, 1e-1, (1 * Gram).Decagrams())
	assert.Equal(t, 1e-2, (1 * Gram).Hectograms())
	assert.Equal(t, 1e-3, (1 * Gram).Kilograms())

	assert.Equal(t, 1e-3, (1 * Kilogram).Megagrams())
	assert.Equal(t, 1e-3, (1 * Megagram).Gigagrams())
	assert.Equal(t, 1e-3, (1 * Gigagram).Teragrams())
	assert.Equal(t, 1e-3, (1 * Teragram).Petagrams())
	assert.Equal(t, 1e-3, (1 * Petagram).Exagrams())
	assert.Equal(t, 1e-3, (1 * Exagram).Zettagrams())
	assert.Equal(t, 1e-3, (1 * Zettagram).Yottagrams())

	// non-SI
	assert.Equal(t, 1e-3, (1 * Kilogram).Tonnes())
	assert.Equal(t, 1e-3, (1 * Tonne).Kilotonnes())
	assert.Equal(t, 1e-3, (1 * Kilotonne).Megatonnes())
	assert.Equal(t, 1e-3, (1 * Megatonne).Gigatonnes())
	assert.Equal(t, 1e-3, (1 * Gigatonne).Teratonnes())
	assert.Equal(t, 1e-3, (1 * Teratonne).Petatonnes())
	assert.Equal(t, 1e-3, (1 * Petatonne).Exatonnes())

	// avoirdupois
	assert.Equal(t, 0.015432358352941428, (1 * Milligram).TroyGrains())
	assert.Equal(t, 0.0022857142857142855, (1 * TroyGrain).AvoirdupoisOunces())
	assert.Equal(t, 0.03657142857142857, (1 * TroyGrain).AvoirdupoisDrams())
	assert.Equal(t, 0.00014285714285714284, (1 * TroyGrain).AvoirdupoisPounds())

	assert.Equal(t, 0.07142857142857142, (1 * AvoirdupoisPound).UsStones())
	assert.Equal(t, 0.058775510204081644, (1 * TroyPound).UkStones())

	// https://en.wikipedia.org/wiki/Quarter_(unit)#Weight
	assert.Equal(t, 0.08818490487395102, (1 * Kilogram).UsQuarters())
	assert.Equal(t, 0.07873652220888486, (1 * Kilogram).UkQuarters())

	// https://en.wikipedia.org/wiki/Hundredweight
	assert.Equal(t, 50.802345439999996, (1 * LongHundredweight).Kilograms())
	assert.Equal(t, 45.35923700000001, (1 * ShortHundredweight).Kilograms())
	assert.Equal(t, 1.0000000000000002, (112 * AvoirdupoisPound).LongHundredweights()) // round error, expected 1e0
	assert.Equal(t, 1e0, (100 * AvoirdupoisPound).ShortHundredweights())

	// troy
	assert.Equal(t, 0.0020833333333333333, (1 * TroyGrain).TroyOunces())
	assert.Equal(t, 0.0026792288807189982, (1 * Gram).TroyPounds())
}
