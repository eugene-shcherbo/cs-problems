package trees

func RightSideView(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	nextNodes := make([]*BinaryTreeNode, 0)
	nextNodes = append(nextNodes, root)

	rightView := make([]int, 0)
	for len(nextNodes) > 0 {
		levelSize := len(nextNodes)
		rightView = append(rightView, nextNodes[levelSize-1].Val)

		for i := 0; i < levelSize; i++ {
			if nextNodes[0].Left != nil {
				nextNodes = append(nextNodes, nextNodes[0].Left)
			}

			if nextNodes[0].Right != nil {
				nextNodes = append(nextNodes, nextNodes[0].Right)
			}

			nextNodes = nextNodes[1:]
		}
	}

	return rightView
}
