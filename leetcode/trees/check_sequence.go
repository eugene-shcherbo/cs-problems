package trees

func IsValidSequence(root *BinaryTreeNode, arr []int) bool {
	return isValidSequence(root, arr, 0)
}

func isValidSequence(root *BinaryTreeNode, arr []int, idx int) bool {
	if root == nil {
		return false
	}

	if idx == len(arr) {
		return false
	}

	if root.Val != arr[idx] {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == arr[idx] && idx == len(arr)-1
	}

	return isValidSequence(root.Left, arr, idx+1) || isValidSequence(root.Right, arr, idx+1)
}
