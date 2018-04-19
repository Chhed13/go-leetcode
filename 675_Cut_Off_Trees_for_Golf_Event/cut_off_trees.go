package _675_Cut_Off_Trees_for_Golf_Event

import (
	"sort"
	"fmt"
)

/* https://leetcode.com/problems/cut-off-trees-for-golf-event/description/

You are asked to cut off trees in a forest for a golf event. The forest is represented as a non-negative 2D map, in this map:

0 represents the obstacle can't be reached.
1 represents the ground can be walked through.
The place with number bigger than 1 represents a tree can be walked through, and this positive number represents the tree's Height.
You are asked to cut off all the trees in this forest in the order of tree's Height - always cut off the tree with lowest Height first. And after cutting, the original place has the tree will become a grass (value 1).

You will start from the point (0, 0) and you should output the minimum steps you need to walk to cut off all the trees. If you can't cut off all the trees, output -1 in that situation.

You are guaranteed that no two trees have the same Height and there is at least one tree needs to be cut off.

Example 1:
Input:
[
 [1,2,3],
 [0,0,4],
 [7,6,5]
]
Output: 6
Example 2:
Input:
[
 [1,2,3],
 [0,0,0],
 [7,6,5]
]
Output: -1
Example 3:
Input:
[
 [2,3,4],
 [0,0,5],
 [8,7,6]
]
Output: 6
Explanation: You started from the point (0,0) and you can cut off the tree in (0,0) directly without walking.
Hint: size of the given matrix will not exceed 50x50.
*/

func cutOffTree(forest [][]int) int {
	m := len(forest)
	n := len(forest[0])

	trees := []Tree{}
	for i, ir := range forest {
		for j, jr := range ir {
			if jr > 1 {
				trees = append(trees, Tree{jr, [2]int{i, j}})
			}
		}
	}

	//fmt.Printf("%+v \n\n", trees)
	sort.Sort(ByHeight(trees))
	//fmt.Printf("%+v \n\n", trees)
	current := [2]int{0, 0}
	sum := 0

	for _, t := range trees {
		step := getStep(forest, current, t.XY, m, n)

		if step < 0 {
			fmt.Printf("return -1. current: %v, next: %v, steps: %v, height: %v\n", current, t.XY, sum, t.Height)
			return -1
		}
		sum += step
		current = t.XY
	}

	return sum
}

func getStep(forest [][]int, start [2]int, finish [2]int, m int, n int) int {

	visited := make([][]bool, m)
	for iv := range visited {
		visited[iv] = make([]bool, n)
	}
	step := 0
	var queue XYQueue
	queue.Push(start)
	visited[start[0]][start[1]] = true

	for ; queue.Size() > 0; {
		s := queue.Size()
		for i := 0; i < s; i++ {
			q := queue.Pull()
			if q[0] == finish[0] && q[1] == finish[1] {
				return step
			}

			for _, next := range [][2]int{
				{q[0] + 1, q[1]},
				{q[0] - 1, q[1]},
				{q[0], q[1] + 1},
				{q[0], q[1] - 1},
			} {
				if next[0] >= 0 && next[0] < m && next[1] >= 0 && next[1] < n &&
					!visited[next[0]][next[1]] && forest[next[0]][next[1]] != 0 {
					queue.Push(next)
					visited[next[0]][next[1]] = true
				}
			}
		}
		step++
	}
	return -1
}

type XYQueue [][2]int

func (xy *XYQueue) Push(n [2]int) {
	*xy = append(*xy, n)
}

func (xy *XYQueue) Pull() [2]int {
	if len(*xy) == 0 {
		return [2]int{-1, -1}
	}
	r := (*xy)[0]
	*xy = (*xy)[1:]
	return r
}

func (xy *XYQueue) Size() int {
	return len(*xy)
}

type Tree struct {
	Height int
	XY     [2]int
}

type ByHeight []Tree

func (h ByHeight) Len() int {
	return len(h)
}

func (h ByHeight) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h ByHeight) Less(i, j int) bool {
	return h[i].Height < h[j].Height
}
