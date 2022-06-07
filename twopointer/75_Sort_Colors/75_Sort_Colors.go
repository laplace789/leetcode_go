package twopointer

//Input: nums = [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	left, right, i := 0, len(nums)-1, 0

	for i <= right && left < right {
		if nums[i] == 2 {
			//2 must put right side of arr
			nums[i] = nums[right]
			nums[right] = 2
			right--
		} else if nums[i] == 0 {
			//1 must put left side of arr
			nums[i] = nums[left]
			nums[left] = 0
			left++
			i++
		} else {
			i++
		}
	}
}
