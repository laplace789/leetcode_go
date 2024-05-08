package array

//Input: nums1 = [1,2,2,1], nums2 = [2,2]
//Output: [2]

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || nums1 == nil || len(nums2) == 0 || nums2 == nil {
		return []int{}
	}

	set := make(map[int]struct{})
	for _, n := range nums1 {
		set[n] = struct{}{}
	}
	var res []int
	//check duplicate
	resSet := make(map[int]interface{})
	for _, n := range nums2 {
		if _, ok := set[n]; ok {
			if _, exist := resSet[n]; !exist {
				res = append(res, n)
				resSet[n] = nil
			}

		}
	}
	return res
}
