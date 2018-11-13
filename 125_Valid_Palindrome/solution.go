package _125_Valid_Palindrome

/*https://leetcode.com/problems/valid-palindrome/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/

/*
- scan from both sides to center, check 1 and len-1 elements
    don't forget to ignore case and trim special chars
*/
import "strings"

func isPalindrome(s string) bool {
    s = normalize(s)
    for i := 0; i < len(s)/2; i++{
        j := len(s)-i-1
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

func normalize(s string) string {
    f := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z':
            return r - 'A' + 'a'
        case r >= 'a' && r <= 'z' || r >= '0' && r <= '9':
            return r
        }
        return -1
    }
    return strings.Map(f, s)
}
