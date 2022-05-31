package linkedlist

import (
	"leetcode_go/model"
)

/**

Input: head = [3,2,0,-4], pos = 1
Output: true

Input: head = [1,1,1,1], pos = -1
Output: false

*/

type ListNode = model.ListNode

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for head != nil {
		if slow.Next != nil {
			slow = slow.Next
		} else {
			return false
		}

		if fast.Next != nil {
			fast = fast.Next
			if fast.Next != nil {
				fast = fast.Next
			} else {
				return false
			}
		} else {
			return false
		}

		if slow == fast {
			return true
		}
	}

	return false
}
