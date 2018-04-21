package _3_Roman_to_Integer

import (
	"testing"
)

type testCase struct {
	tcase  string
	answer int
}

var testCases = []testCase{
	{
		"I", 1,
	},
	{
		"III", 3,
	},
	{
		"IV", 4,
	},
	{
		"IX", 9,
	},
	{
		"XI", 11,
	},
	{
		"XL", 40,
	},
	{
		"L", 50,
	},
	{
		"LVIII", 58,
	},
	{
		"MCMXCIV", 1994,
	},
	{
		"MCM", 1900,
	},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		answer := romanToInt(tc.tcase)
		if answer != tc.answer {
			t.Errorf("expected: %v, got: %v", tc.answer, answer)
		}
	}
}