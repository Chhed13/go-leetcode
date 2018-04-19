package _17_Letter_Combinations_Of_Phone_Number

import "fmt"

/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	comb := []string{}
	for _, t := range translateDigit(rune(digits[0])) {
		if len(digits) > 1 {
			for _, lc := range letterCombinations(digits[1:]) {
				result := string(t) + lc
				comb = append(comb, result)
			}
		} else {
			comb = append(comb, string(t))
		}
	}

	return comb
}

func translateDigit(d rune) string {
	switch d {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	}
	return ""
}
