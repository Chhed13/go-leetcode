package _394_Decode_String

import (
    "testing"
)

type testCase struct {
    tcase  string
    answer string
}

var testCases = []testCase{
    {"3[a]2[bc]", "aaabcbc",},
    {"3[a2[c]]", "accaccacc",},
    {"2[abc]3[cd]ef", "abcabccdcdcdef",},
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := decodeString(tc.tcase)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
