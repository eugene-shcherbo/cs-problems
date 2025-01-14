package stack

import (
	"math"

	"github.com/eugene-shcherbo/cs-problems/leetcode"
)

type ListNode = leetcode.ListNode

func RemoveNodes(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: math.MaxInt, Next: head}
	nodesToKeep := []*ListNode{dummyHead}

	for curr := head; curr != nil; curr = curr.Next {
		for len(nodesToKeep) > 0 && nodesToKeep[len(nodesToKeep)-1].Val < curr.Val {
			nodesToKeep = nodesToKeep[:len(nodesToKeep)-1]
			nodesToKeep[len(nodesToKeep)-1].Next = curr
		}
		nodesToKeep = append(nodesToKeep, curr)
	}

	return dummyHead.Next
}
