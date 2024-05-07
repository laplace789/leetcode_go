package twopointer

import "sort"

/*
Input: nums = [3,1,4,1,5], k = 2
Output: 2
Explanation: There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.

Input: nums = [1,3,1,5,4], k = 0
Output: 1
Explanation: There is one 0-diff pair in the array, (1, 1).
*/

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	start, end := 0, 1
	set := make(map[int]struct{})

	//先排序
	//11345
	sort.Ints(nums)

	for start <= len(nums)-1 && end <= len(nums)-1 {
		diff := nums[end] - nums[start]
		//add 用再去重複的
		add := nums[end] + nums[start]
		if diff < k || start == end {
			//diff比k還小 -> 間距不夠大 -> end 往右
			end++
		} else if diff > k {
			// diff比k還小 -> 間距不夠大 -> start 往右
			start++
		} else {
			//需要去重複
			set[add] = struct{}{}
			start++
			end++
		}
	}

	return len(set)
}
