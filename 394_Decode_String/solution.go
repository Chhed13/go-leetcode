package _394_Decode_String

/*https://leetcode.com/problems/decode-string/

Given an encoded string, return it's decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Examples:

s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
*/

/*
Examples:
* 3[a]2[bc] ->
    * 3 * decodeString(a) + 2 * decodeString(bc)
* 3[a2[c]] ->
    * 3 * decodeString(a2[c])
        * a + 2 * decode(c)
* 2[abc]3[cd]ef ->
    * 2 * decodeString(abc) + 3 * decodeString(cd) + ef

Rules:
* if pattern start with strings -> multiply with content of [] return from decodeString()
* if pattern start with letters -> add it to output as is

* find surraunding [] by counting
*
*/

import "unicode"
import "strconv"
import "strings"

// recursive version
func decodeString(s string) string {
    out := ""
    for i := 0; i < len(s); {
        // multiplier
        bCounter := 0
        mult := 1
        str := ""
        ending := ""
        if unicode.IsDigit(rune(s[i])) {
            startIndex := i
            for unicode.IsDigit(rune(s[i])) {
                i++
            }
            mult, _ = strconv.Atoi(s[startIndex:i])
        }
        if s[i] == '[' {
            i++
            startIndex := i
            bCounter++
            for bCounter > 0 {
                i++
                if s[i] == '[' {
                    bCounter++
                } else if s[i] == ']' {
                    bCounter--
                }
            }
            str = decodeString(s[startIndex:i])
            i++
        }

        if i < len(s) && unicode.IsLetter(rune(s[i])) {
            startIndex := i
            for i < len(s) && unicode.IsLetter(rune(s[i])) {
                i++
            }
            ending = s[startIndex:i]
        }
        out += strings.Repeat(str, mult) + ending
    }
    return out
}

// stack version
func decodeString2(s string) string {
    return s
}
