package _8_String_to_Integer_atoi_

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer int
}

var testCases = []testCase{
    {"   -42", -42},
    {"4193 with words", 4193},
    {"words and 987", 0},
    {"-91283472332", -2147483648},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := myAtoi(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
