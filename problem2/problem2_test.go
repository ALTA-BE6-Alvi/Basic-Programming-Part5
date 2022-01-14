package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaesar(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "def", Caesar(3, "abc"), "Output tidak sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, "cnvc", Caesar(2, "alta"), "Output tidak sesuai")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, "kvdobbkkmknowi", Caesar(10, "alterraacademy"), "Output tidak sesuai")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, "bcdefghijklmnopqrstuvwxyza", Caesar(1, "abcdefghijklmnopqrstuvwxyz"), "Output tidak sesuai")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, "mnopqrstuvwxyzabcdefghijkl", Caesar(1000, "abcdefghijklmnopqrstuvwxyz"), "Output tidak sesuai")
	})
}
