package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxSumSubArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 9, FindaMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}), "Output tidak sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, 7, FindaMaxSumSubArray(2, []int{2, 3, 4, 1, 5}), "Output tidak sesuai")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, 5, FindaMaxSumSubArray(2, []int{2, 1, 4, 1, 1}), "Output tidak sesuai")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, 7, FindaMaxSumSubArray(3, []int{2, 1, 4, 1, 1}), "Output tidak sesuai")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, 8, FindaMaxSumSubArray(4, []int{2, 1, 4, 1, 1}), "Output tidak sesuai")
	})
}
