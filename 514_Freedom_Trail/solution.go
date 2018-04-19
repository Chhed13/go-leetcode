package _514_Freedom_Trail

/*
https://leetcode.com/problems/freedom-trail/description/

In the video game Fallout 4, the quest "Road to Freedom" requires players to reach a metal dial called the "Freedom Trail Ring", and use the dial to spell a specific keyword in order to open the door.

Given a string ring, which represents the code engraved on the outer ring and another string key, which represents the keyword needs to be spelled. You need to find the minimum number of steps in order to spell all the characters in the keyword.

Initially, the first character of the ring is aligned at 12:00 direction. You need to spell all the characters in the string key one by one by rotating the ring clockwise or anticlockwise to make each character of the string key aligned at 12:00 direction and then by pressing the center button.
At the stage of rotating the ring to spell the key character key[i]:
You can rotate the ring clockwise or anticlockwise one place, which counts as 1 step. The final purpose of the rotation is to align one of the string ring's characters at the 12:00 direction, where this character must equal to the character key[i].
If the character key[i] has been aligned at the 12:00 direction, you need to press the center button to spell, which also counts as 1 step. After the pressing, you could begin to spell the next character in the key (next stage), otherwise, you've finished all the spelling.

https://leetcode.com/static/images/problemset/ring.jpg

Example:
Input: ring = "godding", key = "gd"
Output: 4
Explanation:
 For the first key character 'g', since it is already in place, we just need 1 step to spell this character.
 For the second key character 'd', we need to rotate the ring "godding" anticlockwise by two steps to make it become "ddinggo".
 Also, we need 1 more step for spelling.
 So the final output is 4.


Note:
Length of both ring and key will be in range 1 to 100.
There are only lowercase letters in both strings and might be some duplcate characters in both strings.
It's guaranteed that string key could always be spelled by rotating the string ring.
*/

import "fmt"

func findRotateSteps(ring string, key string) int {
	// pre fill
	ringMap := map[rune][]int {}
	for k,v := range ring {
		ringMap[v] = append(ringMap[v], k)
	}

	current := []int{0} //default state
	for ki, k := range key { //move by key letters
		prev := current
		current = make([]int, len(ringMap[k]))
		for ri, r := range ringMap[k] { //iterate on current key letter possibilities on ring
			current[ri] = 1005000 // max value
			for pi, p := range prev { //match variants with previous letter possibilities
				prevPosition := 0
				if ki > 0 {
					prevPosition = ringMap[rune(key[ki-1])][pi] //find position of previous letter on ring
				}

				newR := min(r,prevPosition,len(ring))
				if current[ri] > newR + p {
					current[ri] = newR + p
				}
			}
		}
	}

	result := 1005000
	fmt.Print("posibilities: ")
	for _, c := range current {
		fmt.Print(c, " | ")
		if c < result {
			result = c
		}
	}

	fmt.Printf(" result: %d \n", result)

	return result + len(key)
}

func min (current, newVal, ringSize int)  int {
	//straight
	var t1, t2 int
	if current < newVal {
		t2 = (newVal - current)
	} else {
		t2 = (current - newVal)
	}
	//overlap
	t1 = ringSize - t2
	if t1 < t2 {
		return t1
	} else {
		return t2
	}
}
