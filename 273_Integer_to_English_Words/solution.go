package _273_Integer_to_English_Words

import "strconv"

/*
https://leetcode.com/problems/integer-to-english-words/description/

Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 231 - 1.

For example,
123 -> "One Hundred Twenty Three"
12345 -> "Twelve Thousand Three Hundred Forty Five"
1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
*/

// 1 2 3
// every 3rd digit - allways Diggit + dimention like thousand
// every 2nd digit - tens, execept 1 - it's merged with 1st digit

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	strNum := strconv.Itoa(num)
	result := ""
	triplet := 1
	dimention := 0

	var d1, d2, d3 string
	for i := len(strNum) - 1; i >= 0; i-- {
		switch triplet {
		case 1:
			d3 = getDigit(strNum[i])
			triplet++
		case 2:
			if strNum[i] == '1' {
				d2 = getSecondTen(strNum[i+1])
				d3 = ""
			} else {
				d2 = getTens(strNum[i])
			}
			triplet++
		case 3:
			d1 = getDigit(strNum[i])
			triplet = 1
			if d1 != "" || d2 != "" || d3 != "" {
				result = add(result, d1, d2, d3, dimention)
			}
			d1 = ""
			d2 = ""
			d3 = ""
			dimention++
		}
	}
	if d1 != "" || d2 != "" || d3 != "" {
		result = add(result, d1, d2, d3, dimention)
	}

	return result
}

func add(result string, d1, d2, d3 string, dimention int) string {
	// if result != "" { result = " " + result}
	d := getDimention(dimention)
	if d != "" && result != "" {
		result = " " + result
	}
	result = d + result
	if d3 != "" {
		isEmpty := result == ""
		if !isEmpty {
			result = " " + result
		}
		result = d3 + result
	}
	if d2 != "" {
		isEmpty := result == ""
		if !isEmpty {
			result = " " + result
		}
		result = d2 + result
	}
	if d1 != "" {
		isEmpty := result == ""
		if !isEmpty {
			result = " " + result
		}
		result = d1 + " Hundred" + result
	}

	return result
}

func getDimention(d int) string {
	switch d {
	case 0:
		return ""
	case 1:
		return "Thousand"
	case 2:
		return "Million"
	case 3:
		return "Billion"
	case 4:
		return "Trillion"
	}
	return ""
}

func getDigit(d byte) string {
	switch d {
	case '1':
		return "One"
	case '2':
		return "Two"
	case '3':
		return "Three"
	case '4':
		return "Four"
	case '5':
		return "Five"
	case '6':
		return "Six"
	case '7':
		return "Seven"
	case '8':
		return "Eight"
	case '9':
		return "Nine"
	}
	return ""
}

func getSecondTen(d byte) string {
	switch d {
	case '1':
		return "Eleven"
	case '2':
		return "Twelve"
	case '3':
		return "Thirteen"
	case '4':
		return "Fourteen"
	case '5':
		return "Fifteen"
	case '6':
		return "Sixteen"
	case '7':
		return "Seventeen"
	case '8':
		return "Eighteen"
	case '9':
		return "Nineteen"
	case '0':
		return "Ten"
	}
	return ""
}

func getTens(d byte) string {
	switch d {
	case '1':
		return ""
	case '2':
		return "Twenty"
	case '3':
		return "Thirty"
	case '4':
		return "Forty"
	case '5':
		return "Fifty"
	case '6':
		return "Sixty"
	case '7':
		return "Seventy"
	case '8':
		return "Eighty"
	case '9':
		return "Ninety"
	}
	return ""
}
