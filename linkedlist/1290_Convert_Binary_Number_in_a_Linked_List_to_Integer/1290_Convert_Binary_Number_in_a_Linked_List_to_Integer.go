package linkedlist

import "leetcode_go/model"

type ListNode = model.ListNode

func getDecimalValue(head *ListNode) int {
	if head.Next == nil {
		if head.Val == 1 {
			return 1
		} else {
			return 0
		}
	}

	sum := 0
	for head != nil {
		//last bit
		if head.Next == nil {
			if head.Val == 1 {
				sum += 1
			}
			break
		}
		if head.Val == 1 {
			sum = sum*2 + 2
		} else {
			sum = sum * 2
		}
		head = head.Next
	}

	return sum
}

//  |：bitwise OR
func getDecimalValueHigh(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}
