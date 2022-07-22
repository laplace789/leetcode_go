package linkedlist

import "leetcode_go/model"

type ListNode = model.ListNode

//[1,2,3,4]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}

	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	//two pointer
	preslow, slow, fast := dummyHead, head, head

	//fast move n step
	for fast != nil {
		if n <= 0 {
			preslow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preslow.Next = slow.Next

	return dummyHead.Next
}

func removeNthFromEndTwo(head *ListNode, n int) *ListNode {
	//count len
	if n == 0 {
		return head
	}

	if head == nil {
		return nil
	}
	length := 0
	current := head
	for current != nil {
		current = current.Next
		length++
	}

	if n > length {
		return head
	}

	dummyHead := &ListNode{Next: head}
	dummy := dummyHead
	tmp := head
	for i := 0; i < length-n; i++ {
		dummy = dummy.Next
		tmp = tmp.Next
	}

	dummy.Next = tmp.Next

	return dummyHead.Next
}
