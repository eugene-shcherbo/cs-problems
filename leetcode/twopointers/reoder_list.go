package twopointers

import "github.com/eugene-shcherbo/cs-problems/leetcode"

func ReorderList(head *leetcode.ListNode) {
	midNode := getMidNode(head)
	reversedHalf := reverseList(midNode)

	curr := &leetcode.ListNode{Val: -1, Next: nil}

	for l, r := head, reversedHalf; l != midNode; {
		nextL := l.Next
		nextR := r.Next

		curr.Next = l
		curr.Next.Next = r

		l = nextL
		r = nextR
		curr = curr.Next.Next
	}
}

func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
	var curr *leetcode.ListNode

	for head != nil {
		head.Next, curr, head = curr, head, head.Next
	}

	return curr
}

func getMidNode(head *leetcode.ListNode) *leetcode.ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
