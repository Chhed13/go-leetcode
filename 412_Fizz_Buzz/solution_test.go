package _412_Fizz_Buzz

import (
    "testing"
)

type testCase struct {
    tcase  int
    answer []string
}

var testCases = []testCase{
    {
        15,
        []string{
            "1",
            "2",
            "Fizz",
            "4",
            "Buzz",
            "Fizz",
            "7",
            "8",
            "Fizz",
            "Buzz",
            "11",
            "Fizz",
            "13",
            "14",
            "FizzBuzz",
        },
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := fizzBuzz(tc.tcase)
        for i := range answer {
            if answer[i] != tc.answer[i] {
                t.Errorf("expected: %v, got: %v", tc.answer, answer)
            }
        }
    }
}
