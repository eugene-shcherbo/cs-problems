package trees

type nextRightNode struct {
	val               int
	right, left, next *nextRightNode
}

func ConnectBFS(root *nextRightNode) *nextRightNode {
	if root == nil {
		return nil
	}

	nextNodes := make([]*nextRightNode, 0)
	nextNodes = append(nextNodes, root)

	for len(nextNodes) > 0 {
		levelSize := len(nextNodes)

		for i := 0; i < levelSize; i++ {
			if i < levelSize-1 {
				nextNodes[0].next = nextNodes[1]
			}

			if nextNodes[0].left != nil {
				nextNodes = append(nextNodes, nextNodes[0].left)
			}

			if nextNodes[0].right != nil {
				nextNodes = append(nextNodes, nextNodes[0].right)
			}

			nextNodes = nextNodes[1:]
		}
	}

	return root
}

func ConnectRecursive(root *nextRightNode) *nextRightNode {
	connectRecursive(root, nil)
	return root
}

func connectRecursive(root *nextRightNode, prev *nextRightNode) {
	if root == nil {
		return
	}

	connectRecursive(root.right, root.left)

	if prev != nil {
		prev.next = root
		connectRecursive(root.left, prev.right)
	} else {
		connectRecursive(root.left, nil)
	}
}
