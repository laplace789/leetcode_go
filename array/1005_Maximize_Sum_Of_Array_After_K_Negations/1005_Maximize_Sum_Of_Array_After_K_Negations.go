package array

import "sort"

/*
Example 1:

Input: nums = [4,2,3], k = 1
Output: 5
Explanation: Choose index 1 and nums becomes [4,-2,3].
Example 2:

Input: nums = [3,-1,0,2], k = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
Example 3:

Input: nums = [2,-3,-1,5,-4], k = 2
Output: 13
Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].

*/

func largestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	sort.Ints(nums)
	//Input: nums = [3,-1,0,2], k = 3

	// 1.負著先全部轉正值
	// 2. (1)如果 k%2 == 0 -> 直接加總(正值*-1*-1)
	//    (2)如果 k%2 == 1 -> 找最小的的轉負值
	//
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if k > 0 {
				nums[i] = -1 * nums[i]
				k--
			}
		}
	}

	sort.Ints(nums)

	if k%2 == 1 {
		nums[0] = -1 * nums[0]
	}

	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}

	return res
}
