package backtracking

//Input: nums = [1,2,3]
//Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}
	backtracking(nums, 0, cur, &res)
	return res
}

func backtracking(nums []int, start int, cur []int, res *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*res = append(*res, tmp)
	//[1] [2,3]
	for i := start; i < len(nums); i++ {
		//add condition
		cur = append(cur, nums[i])
		backtracking(nums, i+1, cur, res)
		//remove condition
		cur = cur[:len(cur)-1]
	}
}
