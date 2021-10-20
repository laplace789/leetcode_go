package model

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Int2TreeNode(array []int) *TreeNode {

	n := len(array)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: array[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		//empty queue
		queue = queue[1:]

		if i < n && array[i] != NULL {
			node.Left = &TreeNode{Val: array[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && array[i] != NULL {
			node.Right = &TreeNode{Val: array[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root

	return nil
}
