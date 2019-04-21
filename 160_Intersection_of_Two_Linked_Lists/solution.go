package _160_Intersection_of_Two_Linked_Lists

/* https://leetcode.com/problems/intersection-of-two-linked-lists/
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:


begin to intersect at node c1.



Example 1:


Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.


Example 2:


Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.


Example 3:


Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

/*
if intersect -> return node*
if not -> nil
Properties:
the end of the lists is the same
if I reverse both, it will be not possible to return

Brute force idea: make hashtable of pointers to 1 list, and find first entry of the second

Second idea:
Original
1 -> 2 -> 3 -> 4
8 -> 7 --/

First iteration
1 <- 2 <- 3 <- 4
8 -> 7 --/

Third idea: use stacks - one more brute force
fill the stacks
pop both in sync
first non equal element will be the answer
*/

type ListNode struct {
    Val  int
    Next *ListNode
}

type Stack []*ListNode

func (s *Stack) push(n *ListNode) {
    *s = append(*s, n)
}

func (s *Stack) pop() *ListNode {
    l := len(*s)
    if l == 0 {
        return nil
    }
    out := (*s)[l-1]
    *s = (*s)[:l-1]
    return out
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    sA, sB := Stack{}, Stack{}
    for a := headA; a != nil; a = a.Next {
        sA.push(a)
    }
    for b := headB; b != nil; b = b.Next {
        sB.push(b)
    }

    var out *ListNode
    for a, b := sA.pop(), sB.pop(); a != nil && b != nil; a, b = sA.pop(), sB.pop() {
        if a != b {
            return out
        }
        out = a
    }
    return out
}

