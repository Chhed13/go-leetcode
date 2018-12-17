package _118_Pascals_Triangle

/*https://leetcode.com/problems/pascals-triangle/

Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/


func generate(numRows int) [][]int {
    if numRows == 0 {
        return nil
    }
    out := make([][]int,numRows)
    out[0] = []int{1,}
    for i := 1; i < numRows; i++ {
        out[i] = make([]int,i+1)
        out[i][0] = 1
        for j := 1; j < i; j++ {
            out[i][j] = out[i-1][j-1] + out[i-1][j]
        }
        out[i][i] = 1
    }

    return out
}