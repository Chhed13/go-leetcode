package _204_Count_Primes

import (
    "testing"
)

type testCase struct {
    tcase  int
    answer int
}

var testCases = []testCase{
    {
        10, 4,
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := countPrimes(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}