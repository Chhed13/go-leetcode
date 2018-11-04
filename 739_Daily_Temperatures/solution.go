package _739_Daily_Temperatures

/*https://leetcode.com/problems/daily-temperatures/

Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].
*/

//worse, but simple solution
func dailyTemperatures(T []int) []int {
    if len(T) == 0 {
        return []int{}
    }
    if len(T) == 1 {
        return []int{0}
    }

    out := make([]int,len(T))

    for low,high := 0,1; low < len(T); {
        if high >= len(T){
            out[low] = 0
            low++
            high = low+1
        } else if T[high] > T[low] {
            out[low] = high - low
            low++
            high = low+1
        } else {
            high++
        }
    }

    return out
}
