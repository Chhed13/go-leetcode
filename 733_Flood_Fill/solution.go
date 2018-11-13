package _733_Flood_Fill

/*https://leetcode.com/problems/flood-fill/

An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).

Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.

To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.

At the end, return the modified image.

Example 1:
Input:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation:
From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
by a path of the same color as the starting pixel are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected
to the starting pixel.
Note:

The length of image and image[0] will be in the range [1, 50].
The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
The value of each color in image[i][j] and newColor will be an integer in [0, 65535].
*/

/*
* get initial color oldColor
* add starting Poi to the stack
* add used/visited matrix
* for stack not empty
    * check 4d neigbours if their's color is oldColor and they not visited/used - add them to the stack
*
*/

// one-method, 16ms
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }

    s := [][]int{}

    push := func(sr int, sc int) {
        s = append(s, []int{sr, sc,})
    }

    pop := func() (int, int) {
        if len(s) == 0 {
            return -1, -1
        }

        src := s[len(s)-1]
        s = s[:len(s)-1]
        return src[0], src[1]
    }

    used := make([][]bool, len(image))
    for i := range used {
        used[i] = make([]bool, len(image[0]))
    }

    oldColor := image[sr][sc]
    push(sr, sc)
    used[sr][sc] = true
    for len(s) > 0 {
        pr, pc := pop()
        image[pr][pc] = newColor

        if pr < len(image)-1 && !used[pr+1][pc] && image[pr+1][pc] == oldColor {
            push(pr+1, pc)
            used[pr+1][pc] = true
        }
        if pr > 0 && !used[pr-1][pc] && image[pr-1][pc] == oldColor {
            push(pr-1, pc)
            used[pr-1][pc] = true
        }
        if pc < len(image[pr])-1 && !used[pr][pc+1] && image[pr][pc+1] == oldColor {
            push(pr, pc+1)
            used[pr][pc+1] = true
        }
        if pc > 0 && !used[pr][pc-1] && image[pr][pc-1] == oldColor {
            push(pr, pc-1)
            used[pr][pc-1] = true
        }
    }
    return image
}

// OOP version, 16ms
type Poi struct {
    x, y int
}

type Stack []Poi

func (s *Stack) push(p Poi) {
    *s = append(*s, p)
}

func (s *Stack) pop() Poi {
    if len(*s) == 0 {
        return Poi{}
    }
    p := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return p
}

func (s *Stack) length() int {
    return len(*s)
}

func fitImage(p Poi, i [][]int) bool {
    return p.x >= 0 && p.x < len(i) && p.y >= 0 && p.y < len(i[p.x])
}

func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }

    oldColor := image[sr][sc]
    visited := map[Poi]bool{}
    s := Stack{}
    p := Poi{sr, sc}
    s.push(p)
    visited[p] = true

    for s.length() > 0 {
        p := s.pop()
        image[p.x][p.y] = newColor // if it's expensive -> check visited

        for _, pn := range []Poi{
                {p.x + 1, p.y},
                {p.x - 1, p.y},
                {p.x, p.y + 1},
                {p.x, p.y - 1}} {
            if fitImage(pn, image) && !visited[pn] && image[pn.x][pn.y] == oldColor {
                s.push(pn)
                visited[pn] = true
            }
        }
    }
    return image
}
