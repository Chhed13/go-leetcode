package _234_Palindrome_Linked_List

/*https://leetcode.com/problems/palindrome-linked-list/description/

Given a singly linked list, determine if it is a palindrome.

Follow up:
Could you do it in O(n) time and O(1) space?
*/

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    fast, slow := head, head
    var rev *ListNode
    rev = nil

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow, rev, slow.Next = slow.Next, slow, rev
    }

    if fast != nil {
        slow = slow.Next
    }

    for slow != nil {
        if rev.Val == slow.Val {
            rev = rev.Next
            slow = slow.Next
        } else {
            return false
        }
    }

    return true
}
