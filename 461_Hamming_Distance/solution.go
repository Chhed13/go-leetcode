package _461_Hamming_Distance

/* https://leetcode.com/problems/hamming-distance/

The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different.
*/

/* Desctiption
(1 0 0 1)
(1 1 0 0)
___xor___
(0 1 0 1) <- z

for z > 0
	if z%2 == 1
		count ++
	z = z/2
*/

func hammingDistance(x int, y int) int {
    count := 0
    for z := x ^ y; z > 0; z >>= 1 {
        if z%2 == 1 {
            count++
        }
    }
    return count
}