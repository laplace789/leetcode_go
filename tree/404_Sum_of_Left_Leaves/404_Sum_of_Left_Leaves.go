package tree

import (
	"leetcode_go/model"
)

/**
Input: root = [3,9,20,null,null,15,7]
Output: 24

Input: root = [1,2,3,4,5,6]
Output: 4
*/

type TreeNode = model.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	ret := 0
	if root == nil {
		return 0
	}

	//flag is to identify which side 1:left 2:right
	var traverseSum func(root *TreeNode, flag int)
	traverseSum = func(root *TreeNode, flag int) {
		if root.Left != nil {
			traverseSum(root.Left, 1)
		}
		if root.Right != nil {
			traverseSum(root.Right, 2)
		}
		//leaf
		if root.Left == nil && root.Right == nil {
			if flag == 1 {
				ret += root.Val
			}
		}
	}

	traverseSum(root, 0)
	return ret
}
