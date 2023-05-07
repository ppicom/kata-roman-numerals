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
		{
			n:        5,
			expected: "V",
		},
		{
			n:        6,
			expected: "VI",
		},
		{
			n:        7,
			expected: "VII",
		},
		{
			n:        8,
			expected: "VIII",
		},
		{
			n:        9,
			expected: "IX",
		},
		{
			n:        10,
			expected: "X",
		},
		{
			n:        11,
			expected: "XI",
		},
		{
			n:        12,
			expected: "XII",
		},
		{
			n:        13,
			expected: "XIII",
		},
		{
			n:        14,
			expected: "XIV",
		},
		{
			n:        15,
			expected: "XV",
		},
		{
			n:        16,
			expected: "XVI",
		},
		{
			n:        17,
			expected: "XVII",
		},
		{
			n:        18,
			expected: "XVIII",
		},
		{
			n:        19,
			expected: "XIX",
		},
		{
			n:        20,
			expected: "XX",
		},
		{
			n:        21,
			expected: "XXI",
		},
		{
			n:        39,
			expected: "XXXIX",
		},
		{
			n:        40,
			expected: "XL",
		},
		{
			n:        41,
			expected: "XLI",
		},
		{
			n:        42,
			expected: "XLII",
		},
		{
			n:        43,
			expected: "XLIII",
		},
		{
			n:        44,
			expected: "XLIV",
		},
		{
			n:        45,
			expected: "XLV",
		},
		{
			n:        46,
			expected: "XLVI",
		},
		{
			n:        47,
			expected: "XLVII",
		},
		{
			n:        48,
			expected: "XLVIII",
		},
		{
			n:        49,
			expected: "XLIX",
		},
		{
			n:        50,
			expected: "L",
		},
		{
			n:        60,
			expected: "LX",
		},
		{
			n:        75,
			expected: "LXXV",
		},
		{
			n:        90,
			expected: "XC",
		},
		{
			n:        94,
			expected: "XCIV",
		},
		{
			n:        99,
			expected: "XCIX",
		},
		{
			n:        100,
			expected: "C",
		},
		{
			n:        105,
			expected: "CV",
		},
		{
			n:        149,
			expected: "CXLIX",
		},
		{
			n:        150,
			expected: "CL",
		},
		{
			n:        200,
			expected: "CC",
		},
		{
			n:        300,
			expected: "CCC",
		},
		{
			n:        400,
			expected: "CD",
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
