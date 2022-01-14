package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 4, RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9}), "Output tidak sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 6, RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9}), "Output tidak sesuai")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, 2, RemoveDuplicates([]int{2, 2, 2, 11}), "Output tidak sesuai")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, 2, RemoveDuplicates([]int{2, 2, 2, 11}), "Output tidak sesuai")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, 4, RemoveDuplicates([]int{1, 2, 3, 11, 11}), "Output tidak sesuai")
	})
}
