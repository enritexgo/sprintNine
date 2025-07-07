package main

// Пишите тесты в этом файле

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {

	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"size = 0", generateRandomElements(0), 0},
		{"size < 0", generateRandomElements(-71), 0},
		{"random size", generateRandomElements(71), 71},
	}

	for _, test := range tests {
		t.Run("Check generation with "+test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, len(test.data),
				fmt.Sprintf("Expected lenght of size: %d. Get: %d", test.expected, len(test.data)))
		})
	}
}

func TestMaximum(t *testing.T) {

	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"empty slice", []int{}, 0},
		{"slice with one element", []int{71}, 71},
		{"slice with several elements", []int{91, 2183, 0, 842, 52, 777}, 2183},
	}

	for _, test := range tests {
		t.Run("Checking the maximum search in "+test.name, func(t *testing.T) {
			max := maximum(test.data)
			assert.Equal(t, test.expected, max,
				fmt.Sprintf("Expected count of maximum: %d. Get: %d", test.expected, max))
		})
	}
}
