package _344_Reverse_String

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer string
}

var testCases = []testCase{
    {"hello", "olleh",},
    {"A man, a plan, a canal: Panama", "amanaP :lanac a ,nalp a ,nam A",},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := reverseString(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
