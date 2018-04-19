package _357_Count_Numbers_with_Unique_Digits


// https://leetcode.com/problems/count-numbers-with-unique-digits/description/
//Given a non-negative integer n, count all numbers with unique digits, x, where 0 ≤ x < 10n.
//
//Example:
//Given n = 2, return 91. (The answer should be the total numbers in the range of 0 ≤ x < 100, excluding [11,22,33,44,55,66,77,88,99])
//
//Credits:
//Special thanks to @memoryless for adding this problem and creating all test cases.

//1 - 10
//2 - 91


func countNumbersWithUniqueDigits(n int) int {
	if n > 10 {
		n = 10
	}
	var answer int
	for i := 0; i <= n; i++ {
		answer += table(i)
	}

	return answer
}

func table(i int) int {
	switch i {
	case 0: return 1
	case 1: return 9
	case 2: return 9*9
	case 3: return 9*9*8
	case 4: return 9*9*8*7
	case 5: return 9*9*8*7*6
	case 6: return 9*9*8*7*6*5
	case 7: return 9*9*8*7*6*5*4
	case 8: return 9*9*8*7*6*5*4*3
	case 9: return 9*9*8*7*6*5*4*3*2
	case 10: return 9*9*8*7*6*5*4*3*2*1
	}
	return -1
}