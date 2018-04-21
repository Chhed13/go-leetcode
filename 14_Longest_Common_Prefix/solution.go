package _14_Longest_Common_Prefix

/*https://leetcode.com/problems/longest-common-prefix/description/

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	common := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := range common {
			if j >= len(strs[i])  ||	strs[i][j] != common[j] {
				common = common[:j]
				break
			}
		}
	}

	return common
}