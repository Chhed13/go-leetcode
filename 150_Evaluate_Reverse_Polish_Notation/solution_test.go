package _150_Evaluate_Reverse_Polish_Notation

import (
    "testing"
)

type testCase struct {
    tcase  []string
    answer int
}

var testCases = []testCase{
    {[]string{"2", "1", "+", "3", "*",}, 9,},
    {[]string{"4", "13", "5", "/", "+"}, 6,},
    {[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22,},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := evalRPN(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
