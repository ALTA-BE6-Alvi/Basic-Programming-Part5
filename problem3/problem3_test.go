package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayUnique(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []int{2, 4}, ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}), "Output tidak sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, []int{20, 30, 40}, ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}), "Output tidak sesuai")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, []int{7}, ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}), "Output tidak sesuai")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, []int{3}, ArrayUnique([]int{3, 8}, []int{2, 8}), "Output tidak sesuai")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, []int{}, ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}), "Output tidak sesuai")
	})
}
