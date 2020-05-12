package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParty_Other(t *testing.T) {
	assert.Equal(t, ReParty.Other(), ContraParty)
	assert.Equal(t, ContraParty.Other(), ReParty)
	assert.Panics(t, func() { NoParty.Other() })
}
