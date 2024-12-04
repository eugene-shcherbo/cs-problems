package trees

type treeInfo struct {
	size      int
	found     bool
	kSmallest int
}

func KthSmallest(root *TreeNode, k int) int {
	info := kthSmallest(root, k)

	if info.found {
		return info.kSmallest
	}

	return -1
}

func kthSmallest(root *TreeNode, k int) *treeInfo {
	if root == nil {
		return &treeInfo{0, false, -1}
	}

	leftInfo := kthSmallest(root.Left, k)

	if leftInfo.found {
		return leftInfo
	} else if leftInfo.size == k-1 {
		return &treeInfo{-1, true, root.Val}
	} else {
		rightInfo := kthSmallest(root.Right, k-leftInfo.size-1)
		return &treeInfo{leftInfo.size + rightInfo.size + 1, rightInfo.found, rightInfo.kSmallest}
	}
}
