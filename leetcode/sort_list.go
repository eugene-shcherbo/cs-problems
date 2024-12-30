package leetcode

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMid(head)
	tmp := mid.Next
	mid.Next = nil

	left := SortList(head)
	right := SortList(tmp)

	return merge(left, right)
}

func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(one, two *ListNode) *ListNode {
	head := &ListNode{-1, nil}
	curr := head

	for one != nil && two != nil {
		if one.Val < two.Val {
			curr.Next = one
			one = one.Next
		} else {
			curr.Next = two
			two = two.Next
		}
		curr = curr.Next
	}

	for one != nil {
		curr.Next = one
		curr = curr.Next
		one = one.Next
	}

	for two != nil {
		curr.Next = two
		curr = curr.Next
		two = two.Next
	}

	return head.Next
}
