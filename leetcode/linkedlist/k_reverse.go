package linkedlist

func ReverseKGroup(head *ListNode, k int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 0; i < length/k; i++ {
		prev.Next, prev = reverseKNodes(prev.Next, k)
	}

	return dummy.Next
}

func reverseKNodes(head *ListNode, k int) (*ListNode, *ListNode) {
	var newHead, curr *ListNode = nil, head
	for k > 0 {
		next := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = next
		k--
	}
	head.Next = curr

	return newHead, head
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
