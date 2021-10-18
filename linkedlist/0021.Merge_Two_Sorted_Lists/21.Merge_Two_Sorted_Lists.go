package linkedlist

import (
	"leetcode_go/model"
)

type ListNode = model.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

	Input: l1 = [1,2,4], l2 = [1,3,4]
	Output: [1,1,2,3,4,4]

	loop:

	1 -> merge(1,3) -> merge(2,3) -> merge(4,3) -> merge(4,4) -> merge(4,nil)
    1->next  return l1=1  return l1=2   return l2=3   return l1=4   return l2=4
*/

//split to small array use less memory
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}

	return head.Next
}

/*
func main() {
	arr1 := []int{1, 2, 3,4}
	arr2 := []int{1, 3, 4}

	l1 := model.Int2ListNOde(arr1)
	l2 := model.Int2ListNOde(arr2)

	res := mergeTwoLists2(l1, l2)
	resArr := model.ListNode2Int(res)
	fmt.Println(resArr)
}
*/
