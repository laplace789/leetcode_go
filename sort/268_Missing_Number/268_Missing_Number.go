package sort

import "sort"

//Input: nums = [3,0,1]
//Output: 2
//Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
//0 <= nums[i] <= n
//

func missingNumber(nums []int) int {
	total := len(nums)
	if total == 0 {
		return 0
	}
	res := -1

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//[3,0,1] -> [0,1,3]

	for i := 0; i < total; i++ {
		//missing
		if nums[i] != i {
			res = i
			break
		}
	}
	//last num
	if res == -1 {
		res = total
	}

	return res
}

func missingNumberV2(nums []int) int {
	//Arithmetic sequence
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return (1+len(nums))*len(nums)/2 - sum
}
