package twopointer

/*
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

func ContainerWithMostWater(height []int) int {
	if len(height) == 0 {
		return 0
	}

	//雙指針
	max, start, end := 0, 0, len(height)-1

	for start < end {
		//找出高度,選比較矮的
		//長度 = end - start
		width := end - start
		currentHeight := 0
		if height[start] < height[end] {
			currentHeight = height[start]
			start++
		} else {
			currentHeight = height[end]
			end--
		}

		tmp := currentHeight * width
		if tmp > max {
			max = tmp
		}
	}
	return max
}
