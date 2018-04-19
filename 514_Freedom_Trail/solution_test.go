package _514_Freedom_Trail

import (
	"strconv"
	"testing"
)

var testCases = [][3]string{
	{"godding", "gd", "4"},
	{"caotmcaataijjxi", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx", "137"},
	{"pqwcx", "cpqwx", "13"},
	{"abcde", "ade", "6"},
}

func TestSolution(t *testing.T) {
	for _, tc := range testCases {
		r := findRotateSteps(tc[0], tc[1])
		if strconv.Itoa(r) != tc[2] {
			t.Errorf("invalid result, expected %s, got %s", tc[2], r)
		}
	}
}
