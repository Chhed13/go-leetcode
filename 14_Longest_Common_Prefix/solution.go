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
            if j >= len(strs[i]) || strs[i][j] != common[j] {
                common = common[:j]
                break
            }
        }
    }

    return common
}

/* Description:
- take first char from first string
- scan other first elements
- if ok - add it to accumulation out string
- move to next element

O(array_len * avarage_prefix) time
O(avarage_prefix) memory
*/
func longestCommonPrefix2(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    out := []byte{}
    for i := 0; i < len(strs[0]); i++ {
        c := strs[0][i]
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || strs[j][i] != c {
                return string(out)
            }
        }
        out = append(out, c)
    }
    return string(out)
}
