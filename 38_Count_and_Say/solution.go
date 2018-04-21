package _38_Count_and_Say

import "strconv"

/*https://leetcode.com/problems/count-and-say/description/

The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"
*/

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	out := "1"
	for i := 2; i <= n; i++ {
		cur := ""
		c := 1
		for j := 0; j < len(out); j++ {
			if j+1  < len(out) && out[j] == out[j+1]{
				c++
			} else {
				cur += strconv.Itoa(c) + string(out[j])
				c = 1
			}
		}
		out = cur
	}

	return out
}