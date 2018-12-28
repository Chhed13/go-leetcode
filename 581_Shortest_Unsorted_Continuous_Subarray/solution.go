package _581_Shortest_Unsorted_Continuous_Subarray

/*https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.

*/

/* my ugly solution, but 100% better then others, 36ms
- for
    - find violation
    - after that find min for lower | max for higer values
- for
    - find it's insert positions
- diff it the answer
*/

func findUnsortedSubarray(nums []int) int {
    l := len(nums)
    lv, rv := false,false
    var lMin, rMax = nums[l-1], nums[0]
    for i := 0; i< l-1; i++ {
        if lv {
            lMin = min(lMin, nums[i])
        } else {
            lv = nums[i] > nums[i+1]
        }

        if rv {
            rMax = max(rMax, nums[l-1-i])
        } else {
            rv = nums[l-1-i] < nums[l-2-i]
        }
    }

    if !lv || !rv {
        return 0
    }

    li,ri := 0, l
    for i := 0; i < l-1; i++ {
        if lMin < nums[i]{
            li = i
            break
        }
    }
    for i := 0; i < l-1; i++ {
        if rMax > nums[l-1-i] {
            ri = l-i
            break
        }
    }
    return ri - li
}

func min(n1, n2 int) int {
    if n1 < n2 {
        return n1
    }
    return n2
}

func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}

func findUnsortedSubarray2(nums []int) int {
    l := len(nums)
    begin, end := -1,-2
    ma, mi := nums[0], nums[l-1]
    for i := 0 ; i < l; i++ {
        ma = max(ma, nums[i])
        mi = min(mi, nums[l-i-1])

        if nums[i] < ma {
            end = i
        }
        if nums[l-i-1] > mi {
            begin = l-i-1
        }
    }
    return end - begin +1
}
