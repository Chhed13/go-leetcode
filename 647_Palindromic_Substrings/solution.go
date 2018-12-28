package _647_Palindromic_Substrings

/*https://leetcode.com/problems/palindromic-substrings/

Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
Note:
The input string length won't exceed 1000.
*/

func countSubstringsSlow(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        for j := i + 1; j <= len(s); j++ {
            if isPalindrom(s[i:j]) {
                count++
            }
        }
    }
    return count
}

func isPalindrom(s string) bool {
    if len(s) <= 1 {
        return true
    }
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func countSubstringsFaster(s string) int {
    count := 0
    for i := 0; i < len(s); i++ {
        count += countPalindromes(s, i, i) // count odd
        count += countPalindromes(s, i, i+1) // count even
    }
    return count
}

func countPalindromes(s string, begin, end int) int {
    count := 0
    for i, j := begin, end; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
        count++
    }

    return count
}
