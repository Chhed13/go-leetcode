package _412_Fizz_Buzz

/*https://leetcode.com/problems/fizz-buzz/

Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/

import "strconv"

func fizzBuzz(n int) []string {
    out := []string{}
    for i := 1; i <= n; i++ {
        str := ""
        if i%3 == 0 && i%5 == 0  {
            str = "FizzBuzz"
        } else if i%5 == 0 {
            str = "Buzz"
        } else if i%3 == 0 {
            str = "Fizz"
        } else {
            str = strconv.Itoa(i)
        }
        out = append(out,str)
    }
    return out
}