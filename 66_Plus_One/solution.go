package _66_Plus_One

/*https://leetcode.com/problems/plus-one/

66. Plus One
Easy
631
1165


Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

/*
* move from last to first digit
* make +1 (set mark to count) to last <- it's the beginning of sequence
* if adding make digit smaller -> set mark -> count next
* if mark is set end index digit is 0 -> add -1 element to array with digit 1
*/

func plusOne(digits []int) []int {
    count := true

    for i := len(digits)-1; count && i >= 0; i-- {
        prev := digits[i]
        digits[i] = (digits[i]+1)%10
        count = prev > digits[i]
    }

    if count {
        digits = append([]int{1},digits...)
    }

    return digits
}