package array

/*
	Input: nums = [1,1,1,1,1]
	Output: [1,2,3,4,5]
	Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
*/

func runningSum(nums []int) []int {
	sum := 0
	for i := 0;i<len(nums);i++{
		nums[i] += sum
		sum = nums[i]
	}
	return nums
}

//memory less
func runningSum2(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
