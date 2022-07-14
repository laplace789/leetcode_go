package linkedlist

import "leetcode_go/model"

type ListNode = model.ListNode

func middleNode(head *ListNode) *ListNode {
	//bound
	if head.Next == nil {
		return head
	}

	fast := head //move two steps
	slow := head //move one step

	// fast.Next == nil and slow.Next = middle point
	for head != nil {
		if fast.Next != nil {
			fast = fast.Next
			if fast.Next != nil {
				fast = fast.Next
			} else {
				break
			}
		} else {
			break
		}
		//odd cond
		if fast.Next == nil {
			break
		}

		if slow.Next != nil {
			slow = slow.Next
		}
	}
	return slow.Next
}
