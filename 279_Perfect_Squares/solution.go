package _279_Perfect_Squares

/*https://leetcode.com/problems/perfect-squares/

Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/

func numSquares(n int) int {
    var queue []int

    enqueue := func(e int) {
        queue = append(queue, e)
    }

    dequeue := func() int {
        if len(queue) == 0 {
            return -1
        }
        e := queue[0]
        queue = queue[1:]
        return e
    }

    if n <= 0 {
        return 0
    }
    enqueue(n)
    out := 0
    for len(queue) > 0 {
        l := len(queue)
        out++
        for l > 0 {
            nn := dequeue()

            for i := 1; i*i <= nn; i++ {
                if i*i == nn {
                    return out
                }
                enqueue(nn - i*i)
            }
            l--
        }
    }

    return 0

}
