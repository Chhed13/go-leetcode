package _88_Merge_Sorted_Array

import (
    "testing"
)

type testCase struct {
    nums1  []int
    m      int
    nums2  []int
    n      int
    answer []int
}

var testCases = []testCase{
    {
        []int{1, 2, 3, 0, 0, 0}, 3,
        []int{2, 5, 6}, 3,
        []int{1, 2, 2, 3, 5, 6},
    },
}

func TestSolution(t *testing.T) {
    for _, tc := range testCases {

        merge(tc.nums1,tc.m,tc.nums2,tc.n)
        for i := range tc.answer {
            if tc.nums1[i] != tc.answer[i] {
                t.Errorf("expected: %v, got: %v", tc.answer, tc.nums1)
            }
        }
    }
}
