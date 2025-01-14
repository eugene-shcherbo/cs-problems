package linkedlist

import "github.com/eugene-shcherbo/cs-problems/leetcode"

type ListNode = leetcode.ListNode

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
