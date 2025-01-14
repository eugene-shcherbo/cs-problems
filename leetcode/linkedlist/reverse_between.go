package linkedlist

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	curr := dummyHead

	for i := 0; i != left-1; i++ {
		curr = curr.Next
	}

	var prev *ListNode
	oldPrev := curr
	oldLeft := curr.Next
	curr = curr.Next

	for i := 0; i < right-left+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	oldLeft.Next = curr
	oldPrev.Next = prev

	return dummyHead.Next
}
