package main

import (
	"testing"
	"fmt"
)

func Test_RomanNumerals(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{
			n:        1,
			expected: "I",
		},
		{
			n:        2,
			expected: "II",
		},
		{
			n:        3,
			expected: "III",
		},
		{
			n:        4,
			expected: "IV",
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := romanNumerals(tt.n); got != tt.expected {
				t.Errorf("expected %d to be %q, got %q instead", tt.n, tt.expected, got)
			}
		})
	}
}
