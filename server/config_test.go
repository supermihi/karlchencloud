package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadConfig(t *testing.T) {
	c, err := ReadConfig()
	assert.Nil(t, err)
	assert.False(t, c.NoProxy)
}
