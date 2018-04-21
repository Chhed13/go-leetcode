package _20_Valid_Parentheses

/*https://leetcode.com/problems/valid-parentheses/description/

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

func isValid(s string) bool {
	var stack []rune

	push := func(r rune) {
		stack = append(stack, r)
	}

	pop := func() rune{
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return r
	}

	length := func() int{
		return len(stack)
	}

	for _, c := range s {
		switch c {
		case '(':
			push(')')
		case '{':
			push('}')
		case '[':
			push(']')
		default:
			if length() == 0 && c != 0 || length() > 0 && c != pop() {
				return false
			}
		}
	}
	return length() == 0
}
