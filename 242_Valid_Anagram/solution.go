package _242_Valid_Anagram

/*https://leetcode.com/problems/valid-anagram/

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

/* Comment
- simplify:
    check len - if not same -> out;
    make [26]int (time the letter presented) for and scan s,t
    scan s, t at same time, s +1, t -1 on letter
    at the end all counter of array should be 0
O(n) time, O(1) memory
*/
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    d := [26]int{}

    for i := range s {
        d[getIndex(s[i])]++
        d[getIndex(t[i])]--
    }

    for _,v := range d {
        if v != 0 {
            return false
        }
    }
    return true
}

func getIndex(c byte) int {
    return int(c - 'a')
}