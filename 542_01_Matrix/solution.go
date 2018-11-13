package _542_01_Matrix

/* https://leetcode.com/problems/01-matrix/

Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.
Example 1:
Input:

0 0 0
0 1 0
0 0 0
Output:
0 0 0
0 1 0
0 0 0
Example 2:
Input:

0 0 0
0 1 0
1 1 1
Output:
0 0 0
0 1 0
1 2 1
Note:
The number of elements of the given matrix will not exceed 10,000.
There are at least one 0 in the given matrix.
The cells are adjacent in only four directions: up, down, left and right.
*/


/*
* init output matrix by: set all possible 0 by scannin input matrix, others - to -1
* take "0" cells of output matix as starting points for bfs solution
    * add them to the Queue

* pulling Queue while not empty
    * check neighbours, if -1 then make it (current cell count + 1) and add them to Queue
*/

func updateMatrix(matrix [][]int) [][]int {
    if len(matrix) == 0  {
        return matrix
    }

    q := [][]int{}

    push := func(x,y int) {
        q = append(q,[]int{x,y})
    }

    pull := func() (int,int) {
        if len(q) == 0 {
            return -1,-1
        }

        out := q[0]
        q = q[1:]
        return out[0],out[1]
    }

    out := make([][]int,len(matrix))
    for i := range matrix {
        out[i] = make([]int,len(matrix[i]))
        for j := range matrix[i] {
            if matrix[i][j] == 0 {
                out[i][j] = 0
                push(i,j)
            } else {
                out[i][j] = -1
            }
        }
    }

    for len(q) > 0 {
        i,j := pull()
        if i + 1 < len(out) && out[i+1][j] == -1 {
            out[i+1][j] = out[i][j] + 1
            push(i+1,j)
        }
        if i > 0 && out[i-1][j] == -1 {
            out[i-1][j] = out[i][j] + 1
            push(i-1,j)
        }
        if j + 1 < len(out[i]) && out[i][j+1] == -1 {
            out[i][j+1] = out[i][j] + 1
            push(i,j+1)
        }
        if j > 0 && out[i][j-1] == -1 {
            out[i][j-1] = out[i][j] + 1
            push(i,j-1)
        }
    }

    return out
}