package _350_Intersection_of_Two_Arrays_II

/*https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersect(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	var long []int
	if len(nums1) > len(nums2) {
		for _, n2 := range nums2 {
			cache[n2]++
		}
		long = nums1
	} else {
		for _, n1 := range nums1 {
			cache[n1]++
		}
		long = nums2
	}

	var out []int
	for _,l := range long {
		if cache[l] > 0 {
			out = append(out,l)
			cache[l]--
		}
	}
	return out
}
