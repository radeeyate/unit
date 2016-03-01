package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountOfSubstance(t *testing.T) {

	// SI
	assert.Equal(t, 1e0, (1 * Mole).Moles())
}
