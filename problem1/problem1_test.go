package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	t.Run("Case 1:", func(t *testing.T) {
		assert.Equal(t, "AKA", Compare("AKA", "AKASHI"), "Hasil output tidak sesuai")
	})
	t.Run("Case 2:", func(t *testing.T) {
		assert.Equal(t, "KANG", Compare("KANGOORO", "KANG"), "Hasil output tidak sesuai")
	})
	t.Run("Case 3:", func(t *testing.T) {
		assert.Equal(t, "KI", Compare("KI", "KIJANG"), "Hasil output tidak sesuai")
	})
	t.Run("Case 4:", func(t *testing.T) {
		assert.Equal(t, "KUPU", Compare("KUPU-KUPU", "KUPU"), "Hasil output tidak sesuai")
	})
	t.Run("Case 5:", func(t *testing.T) {
		assert.Equal(t, "ILA", Compare("ILALANG", "ILA"), "Hasil output tidak sesuai")
	})
}
