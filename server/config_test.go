package server

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_ = os.Setenv("CONSTANT_TABLE_ID", "12345")
	_ = os.Setenv("NO_PROXY", "true")
	c, err := ReadConfig()
	assert.Nil(t, err)
	assert.True(t, c.NoProxy)
	assert.Equal(t, "12345", c.Room.ConstantTableId)
}
