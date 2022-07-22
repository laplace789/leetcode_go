package linkedlist

import "leetcode_go/model"

type ListNode = model.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	dummy := dummyHead
	for dummy != nil && dummy.Next != nil {
		a := dummy.Next
		b := dummy.Next.Next
		a.Next = b.Next
		dummy.Next = b
		dummy.Next.Next = a
		dummy = dummy.Next.Next
	}
	return dummy.Next
}
