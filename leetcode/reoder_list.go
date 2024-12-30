package leetcode

func ReorderList(head *ListNode) {
	midNode := getMidNode(head)
	reversedHalf := reverseList(midNode)

	curr := &ListNode{-1, nil}

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

func reverseList(head *ListNode) *ListNode {
	var curr *ListNode

	for head != nil {
		head.Next, curr, head = curr, head, head.Next
	}

	return curr
}

func getMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
