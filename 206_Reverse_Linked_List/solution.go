package _206_Reverse_Linked_List

/*https://leetcode.com/problems/reverse-linked-list/description/

Reverse a singly linked list.

click to show more hints.

Hint:
A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	prev = nil

	for head.Next != nil {
		next := head.Next

		head.Next = prev

		prev = head

		head = next
	}

	head.Next = prev

	return head
}
