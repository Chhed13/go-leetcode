package _69_Sqrt_x

/*https://leetcode.com/problems/sqrtx/description/

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	min, max, mid := 1, x, 0
	for {
		mid := min + (max-min)/2
		if mid*mid > x {
			max = mid - 1
		} else {
			if (mid + 1)*(mid + 1) > x {
				return mid
			}
			min = mid + 1
		}
	}
	return mid
}
