package _283_Move_Zeroes

/*https://leetcode.com/problems/move-zeroes/

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

// O(nums) time, O(1) memory
func moveZeroes(nums []int)  {
    if len(nums) <= 1 {
        return
    }

    c := 0
    for i := range nums {
        if nums[i] == 0 {
            c++
        } else {
            nums[i - c] = nums[i]
        }
    }

    for i := len(nums) - c; i < len(nums); i++ {
        nums[i] = 0
    }
}