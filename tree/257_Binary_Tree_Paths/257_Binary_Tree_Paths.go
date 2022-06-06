package main

import (
	"fmt"
	"leetcode_go/model"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = model.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	var ret []string

	if root == nil {
		return ret
	}

	var traverse func(root *TreeNode, path string)
	traverse = func(root *TreeNode, path string) {
		path += strconv.Itoa(root.Val)
		if root.Left != nil {
			traverse(root.Left, path+"->")
		}
		if root.Right != nil {
			traverse(root.Right, path+"->")
		}
		//leaf
		if root.Left == nil && root.Right == nil {
			ret = append(ret, path)
		}
	}
	traverse(root, "")
	return ret
}

func main() {
	root := model.Int2TreeNode([]int{1, 2, 3, model.NULL, 5})
	res := binaryTreePaths(root)
	fmt.Println(res)
}
