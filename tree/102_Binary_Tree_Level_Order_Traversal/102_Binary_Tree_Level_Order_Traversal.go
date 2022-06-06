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

//Input: root = [3,9,20,null,null,15,7]
//Output: [[3],[9,20],[15,7]]

type TreeNode = model.TreeNode

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	level := make([]int, 0)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, level)
		level = []int{}
		queue = queue[n:]
	}
	return res
}
