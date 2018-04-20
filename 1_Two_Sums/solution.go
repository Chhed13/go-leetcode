package _1_Two_Sums

/* https://leetcode.com/problems/two-sum/description/

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	result := make([]int,2)
	index := make(map[int]int,len(nums))
	for i,n := range nums {
		if j,ok := index[target-n]; ok {
			result[1] = i
			result[0] = j
			break
		}
		index[n] = i
	}

	return result
}