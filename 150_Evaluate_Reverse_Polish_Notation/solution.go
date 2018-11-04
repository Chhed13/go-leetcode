package _150_Evaluate_Reverse_Polish_Notation

/*https://leetcode.com/problems/evaluate-reverse-polish-notation/

Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
Example 1:

Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
*/

import "strconv"

func isOperand(t string) bool {
    return t == "+" || t == "-" || t == "/" || t == "*"
}

// forward variant
func evalRPN(tokens []string) int {
    stack := []int{}

    push := func(t int) {
        stack = append(stack, t)
    }

    pop := func() int {
        if len(stack) == 0 {
            return 0
        }
        t := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        return t
    }

    for _, t := range tokens {
        if isOperand(t) {
            o2 := pop()
            o1 := pop()
            switch t {
            case "*":
                push(o1 * o2)
            case "/":
                push(o1 / o2)
            case "+":
                push(o1 + o2)
            case "-":
                push(o1 - o2)
            }
        } else {
            i, _ := strconv.Atoi(t)
            push(i)
        }
    }
    return pop()
}
