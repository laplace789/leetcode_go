package linkedlist

import (
	"leetcode_go/model"
)

/*

Input: head = [4,5,1,9], node = 5
Output: [4,1,9]

*/

type ListNode = model.ListNode

func deleteNode(node *ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			return
		}
		deleteNode(node.Next)
		return
	}
}
