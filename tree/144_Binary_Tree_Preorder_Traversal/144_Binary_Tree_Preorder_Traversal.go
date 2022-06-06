package tree

import (
	"leetcode_go/model"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Input: root = [1,null,2,3]
//Output: [1,2,3]

type TreeNode = model.TreeNode

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var traverse func(root *TreeNode, res *[]int)
	traverse = func(root *TreeNode, ans *[]int) {
		*ans = append(*ans, root.Val)
		if root.Left != nil {
			traverse(root.Left, ans)
		}
		if root.Right != nil {
			traverse(root.Right, ans)
		}
	}
	traverse(root, &ans)
	return ans
}

func preorderTraversalHighLevel(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	list := []int{}
	list = append(list, root.Val)
	if root.Left != nil {
		list = append(list, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
