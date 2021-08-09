package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCoinbase(t *testing.T) {
	assert.True(t, ContainsString([]string{"a", "b"}, "a"))
	assert.False(t, ContainsString([]string{}, "a"))
	assert.False(t, ContainsString([]string{"a", "b"}, "c"))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2))
	assert.Equal(t, 1, Min(2, 1))
	assert.Equal(t, -1, Min(-1, 2))
	assert.Equal(t, 0, Min(0, 0))
}
