package trees

import (
	"math"

	"github.com/eugene-shcherbo/cs-problems/utils"
)

func MaxPathSum(root *BinaryTreeNode) int {
	_, totalMax := maxPathSumHelper(root)
	return totalMax
}

func maxPathSumHelper(root *BinaryTreeNode) (int, int) {
	if root == nil {
		return math.MinInt, math.MinInt
	}

	maxToMeLeft, maxFromLeft := maxPathSumHelper(root.Left)
	maxToMeRight, maxFromRight := maxPathSumHelper(root.Right)

	maxToMeLeft = utils.Max(maxToMeLeft, 0)
	maxToMeRight = utils.Max(maxToMeRight, 0)

	maxSoFar := utils.Max(maxToMeLeft+maxToMeRight+root.Val, utils.Max(maxFromLeft, maxFromRight))
	maxOnPath := utils.Max(maxToMeLeft+root.Val, maxToMeRight+root.Val)

	return maxOnPath, maxSoFar
}
