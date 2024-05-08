package array

/*

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

*/

func searchInsert(nums []int, target int) int {
	//init data
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + (end - start)) >> 1
		if nums[mid] >= target {
			//start - mid 的區間
			end = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			start = mid + 1
		}
	}
	return 0
}
