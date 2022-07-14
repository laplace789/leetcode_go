package tree

import "leetcode_go/model"

type TreeNode = model.TreeNode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	//leaf
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		} else {
			return false
		}
	}
	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}
