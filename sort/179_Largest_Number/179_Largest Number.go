package sort

import (
	"sort"
	"strconv"
)

//Input: nums = [10,2]
//Output: "210"

func largestNumber(nums []int) string {
	res := ""
	if len(nums) == 0 || nums == nil {
		return res
	}

	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		// "30" + "3" = "303"
		//  "3"+"30"  = "330"
		if first+second > second+first {
			return true
		}
		return false
	})

	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}

	//[0,0,0,0] -> "0"
	if nums[0] == 0 {
		return "0"
	}

	return res
}
