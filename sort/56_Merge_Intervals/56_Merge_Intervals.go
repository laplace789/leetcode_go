package sort

import "sort"

//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 || intervals == nil {
		return res
	}

	// sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	//add first to res
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		curStart := intervals[i][0]
		curEnd := intervals[i][1]

		preEnd := intervals[i-1][1]

		if preEnd >= curStart {
			intervals[i-1][1] = max(preEnd, curEnd)
			//replace intervals[i] to compare next time
			intervals[i] = intervals[i-1]
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// [[1,3],[2,6],[8,10],[15,18]]
// [[1,999],[2,6],[8,10],[15,18]]
