package linkedlist

import (
	"leetcode_go/model"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
*/

type ListNode = model.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//head virtual head
	//carry is for next digit
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}

	return head.Next
}

/*
func main(){
	arr1 := []int{2,4,3}
	arr2 := []int{5,6,4}

	l1 := model.Int2ListNOde(arr1)
	l2 := model.Int2ListNOde(arr2)

	resNode := addTwoNumbers(l1,l2)
	fmt.Println(model.ListNode2Int(resNode))
}*/
