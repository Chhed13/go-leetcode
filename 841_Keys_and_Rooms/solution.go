package _841_Keys_and_Rooms

/*https://leetcode.com/problems/keys-and-rooms/

There are N rooms and you start in room 0.  Each room has a distinct number in 0, 1, 2, ..., N-1, and each room may have some keys to access the next room.

Formally, each room i has a list of keys rooms[i], and each key rooms[i][j] is an integer in [0, 1, ..., N-1] where N = rooms.length.  A key rooms[i][j] = v opens the room with number v.

Initially, all the rooms start locked (except for room 0).

You can walk back and forth between rooms freely.

Return true if and only if you can enter every room.

Example 1:

Input: [[1],[2],[3],[]]
Output: true
Explanation:
We start in room 0, and pick up key 1.
We then go to room 1, and pick up key 2.
We then go to room 2, and pick up key 3.
We then go to room 3.  Since we were able to go to every room, we return true.
Example 2:

Input: [[1,3],[3,0,1],[2],[0]]
Output: false
Explanation: We can't enter the room with number 2.
Note:

1 <= rooms.length <= 1000
0 <= rooms[i].length <= 1000
The number of keys in all rooms combined is at most 3000.
*/


/* Description
* track by map[]bool visited - if enter, set true

* enter first room (assume that we have key from 0)
* add key 0 to stack
* for stack not empty
    * pop key, check visited[key], if not -> enter the room
    * track -> set visited[room] = true
    * take keys -> add to the stack

* scan the visited list if all true -> return true
*/


func canVisitAllRooms(rooms [][]int) bool {
    visited := map[int]bool{}

    s := []int{}
    push := func(i int) {
        s = append(s,i)
    }
    pop := func() int {
        if len(s) <= 0 {
            return -1
        }

        r := s[len(s)-1]
        s = s[:len(s)-1]
        return r
    }

    push(0)
    for len(s) > 0 {
        k := pop()
        if visited[k] {
            continue
        }
        visited[k] = true
        for _,i := range rooms[k] {
            if !visited[i] {
                push(i)
            }
        }
    }

    return len(visited) == len(rooms)
}