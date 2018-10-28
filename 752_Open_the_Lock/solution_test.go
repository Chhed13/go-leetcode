package _752_Open_the_Lock

import (
    "testing"
)

type testCase struct {
    target   string
    deadends []string
    answer   int
}

var testCases = []testCase{
    {
        "0202", []string{"0201", "0101", "0102", "1212", "2002"}, 6,
    }, {
        "0009", []string{"8888"}, 1,
    }, {
        "8888", []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, -1,
    }, {
        "8888", []string{"0000"}, -1,
    }, {
        "2010", []string{"2110", "0202", "1222", "2221", "1010"}, 3,
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {
        answer := openLock(tc.deadends, tc.target)
        t.Logf("answer: %v", answer)
        if answer != tc.answer {
            t.Errorf("expected: %v, got: %v", tc.answer, answer)
        }
    }
}
