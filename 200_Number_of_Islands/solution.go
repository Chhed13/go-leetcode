package _200_Number_of_Islands

/*https://leetcode.com/problems/number-of-islands/
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands.
An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
    You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3

*/

func numIslands(grid [][]byte) int {
    var queue [][]int

    enqueue := func(x int, y int) {
        queue = append(queue, []int{x, y})
    }

    dequeue := func() (int, int) {
        if len(queue) == 0 {
            return -1, -1
        }
        r := queue[0]
        queue = queue[1:]
        return r[0], r[1]
    }

    if grid == nil || len(grid) <= 0 {
        return 0
    }

    used := make([][]bool, len(grid))
    for i := range used {
        used[i] = make([]bool, len(grid[0]))
    }
    islands := 0

    for i := range grid {
        for j := range grid[i] {
            if used[i][j] || grid[i][j] == 0 {
                continue
            }

            enqueue(i, j)
            used[i][j] = true
            islands++

            for ; len(queue) > 0; {
                x, y := dequeue()
                if x < len(grid)-1 && grid[x+1][y] > 0 && !used[x+1][y] {
                    enqueue(x+1, y)
                    used[x+1][y] = true
                }
                if x > 0 && grid[x-1][y] > 0 && !used[x-1][y] {
                    enqueue(x-1, y)
                    used[x-1][y] = true
                }
                if y < len(grid[x])-1 && grid[x][y+1] > 0 && !used[x][y+1] {
                    enqueue(x, y+1)
                    used[x][y+1] = true
                }
                if y > len(grid[x]) && grid[x][y-1] > 0 && !used[x][y+1] {
                    enqueue(x, y-1)
                    used[x][y+1] = true
                }
            }

        }
    }
    return islands
}
