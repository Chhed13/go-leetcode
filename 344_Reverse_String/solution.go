package _344_Reverse_String

/*https://leetcode.com/problems/reverse-string/

Write a function that takes a string as input and returns the string reversed.

Example 1:

Input: "hello"
Output: "olleh"
Example 2:

Input: "A man, a plan, a canal: Panama"
Output: "amanaP :lanac a ,nalp a ,nam A"
*/

func reverseString(s string) string {
    out := []rune(s)
    for j, i := len(s)-1, 0; i < len(out)/2; i++ {
        out[i], out[j] = out[j], out[i]
        j--
    }
    return string(out)
}