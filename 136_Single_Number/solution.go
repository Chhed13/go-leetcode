package _136_Single_Number

/*https://leetcode.com/problems/single-number/description/

Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

// crazy un-optimized solution
func singleNumber(nums []int) int {
    buf := map[int]bool{}
    for _, n := range nums {
        if buf[n] {
            delete(buf, n)
        } else {
            buf[n] = true
        }
    }

    var out int
    for i := range buf {
        out = i
    }
    return out
}

func singleNumber2(nums []int) int {
    out := 0
    for _,i := range nums {
        out ^= i
    }
    return out
}