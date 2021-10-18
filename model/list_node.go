package model

type ListNode struct {
	Val  int
	Next *ListNode
}

func Int2ListNOde(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	headNext := head
	for _, val := range arr {
		headNext.Next = &ListNode{Val: val}
		headNext = headNext.Next
	}

	return head.Next
}

func ListNode2Int(node *ListNode) []int {
	if node == nil {
		return []int{}
	}

	arr := make([]int, 0)
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}

	return arr
}

func Int2ListNOdeWithCycle(arr []int, cycleHeadIndex int) *ListNode {
	var cycleHead *ListNode
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	headNext := head
	for index, val := range arr {
		headNext.Next = &ListNode{Val: val}
		headNext = headNext.Next
		if index == cycleHeadIndex {
			cycleHead = headNext
		}
		//last index form a cycle
		if index == len(arr)-1 {
			headNext.Next = cycleHead
		}
	}
	return head.Next
}
