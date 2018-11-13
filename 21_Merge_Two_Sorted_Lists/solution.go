package _21_Merge_Two_Sorted_Lists

import "fmt"

/*https://leetcode.com/problems/merge-two-sorted-lists/description/

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

//Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, cur, prev *ListNode

    for l1 != nil || l2 != nil {
        if l1 == nil {
            cur = l2
            l2 = l2.Next
        } else if l2 == nil {
            cur = l1
            l1 = l1.Next
        } else if l1.Val < l2.Val {
            cur = l1
            l1 = l1.Next
        } else {
            cur = l2
            l2 = l2.Next
        }

        fmt.Println(cur.Val)

        if head == nil {
            head = cur
        }

        if prev != nil {
            prev.Next = cur
        }
        prev = cur
    }

    return head
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
    var r *ListNode
    if l1 == nil {
        r = l2
    } else if l2 == nil {
        r = l1
    } else if l1.Val > l2.Val {
        l2.Next = mergeTwoLists2(l1, l2.Next)
        r = l2
    } else {
        l1.Next = mergeTwoLists2(l1.Next, l2)
        r = l1
    }
    return r
}
