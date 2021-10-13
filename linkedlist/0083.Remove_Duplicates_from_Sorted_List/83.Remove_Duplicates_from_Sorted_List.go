package main

import (
	"fmt"
	"leetcode_go/model"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = model.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	arr1 := []int{1, 1, 2, 4}
	l1 := model.Int2ListNOde(arr1)
	res := deleteDuplicates(l1)
	resArr := model.ListNode2Int(res)
	fmt.Println(resArr)
}
