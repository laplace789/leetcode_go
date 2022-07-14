package tree

import "leetcode_go/model"

type TreeNode = model.TreeNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	return transversal(root, low, high)
}

func transversal(node *TreeNode, low int, high int) int {
	if node == nil {
		return 0
	}
	switch {
	case node.Val > low:
		return transversal(node.Left, low, high)
	case node.Val < high:
		return transversal(node.Right, low, high)
	default:
		return transversal(node.Left, low, high) + node.Val + transversal(node.Right, low, high)
	}
}
