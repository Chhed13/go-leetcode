package _387_First_Unique_Character_in_a_String

/*https://leetcode.com/problems/first-unique-character-in-a-string

Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/


/*
- map[char]int - index count existence of chars in string, on next round return first char that count < 2
O(n+n) time, O(n <= 26) memory
- with []int implementation can be faster
*/

func firstUniqChar(s string) int {
    m := map[rune]int{}
    for _,c := range s {
        m[c]++
    }

    for i,c := range s {
        if m[c] < 2 {
            return i
        }
    }
    return -1
}