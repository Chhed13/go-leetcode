package _88_Merge_Sorted_Array

/* https://leetcode.com/problems/merge-sorted-array

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/


/*
O(n) time and O(1) space
Description: iterate from end of numbers of arrays (n-1 and m-1) backwards and fill the array nums1 from last position
*/

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i,j := m-1,n-1
    for ; i >= 0 && j >= 0; {
        if nums1[i] > nums2[j] {
            nums1[i+j+1] = nums1[i]
            i--
        } else {
            nums1[i+j+1] = nums2[j]
            j--
        }
    }
    for ; j >= 0; j-- {
        nums1[j] = nums2[j]
    }
}