package _268_Missing_Number

/*https://leetcode.com/problems/missing-number/

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
*/

/*
1. get first index and move arround array[prev_value] setting all checked elements to -1. Scan one more time - get last non(-1) number. Space O(1), time O(2n). Sideffect: original array missed
2. Calculate sum of squance [0:n] = median of [0:n] * number of expected elements (len+1), subtract in cycle, what's left - is missing number.
*/
func missingNumber(nums []int) int {
    sum := (len(nums)+1)*(0+len(nums))/2
    for i := range nums {
        sum -= nums[i]
    }
    return sum
}