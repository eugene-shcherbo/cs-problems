package trees

import "github.com/eugene-shcherbo/cs-problems/leetcode"

type BinaryTreeNode = leetcode.BinaryTreeNode

func LevelOrderBFS(root *BinaryTreeNode) [][]int {
	if root == nil {
		return nil
	}

	levels := make([][]int, 0)

	nextNodes := make([]*BinaryTreeNode, 0)
	nextNodes = append(nextNodes, root)

	for len(nextNodes) > 0 {
		levels = append(levels, make([]int, 0))
		numOfNodesOnCurrLevel := len(nextNodes)

		for i := 0; i < numOfNodesOnCurrLevel; i++ {
			levels[len(levels)-1] = append(levels[len(levels)-1], nextNodes[0].Val)

			if nextNodes[0].Left != nil {
				nextNodes = append(nextNodes, nextNodes[0].Left)
			}

			if nextNodes[0].Right != nil {
				nextNodes = append(nextNodes, nextNodes[0].Right)
			}

			nextNodes = nextNodes[1:]
		}
	}

	return levels
}

func LevelOrderRecursive(root *BinaryTreeNode) [][]int {
	levels := make([][]int, 0)
	levelOrderRecursiveImpl(root, 0, &levels)

	return levels
}

func levelOrderRecursiveImpl(root *BinaryTreeNode, level int, levels *[][]int) {
	if root == nil {
		return
	}

	if level >= len(*levels) {
		*levels = append(*levels, make([]int, 0))
	}

	(*levels)[level] = append((*levels)[level], root.Val)

	levelOrderRecursiveImpl(root.Left, level+1, levels)
	levelOrderRecursiveImpl(root.Right, level+1, levels)
}
