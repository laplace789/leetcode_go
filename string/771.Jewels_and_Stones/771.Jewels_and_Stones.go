package string

/*

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3


Input: jewels = "z", stones = "ZZ"
Output: 0

*/
func numJewelsInStones(jewels string, stones string) int {
	count := 0

	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if string(stones[j]) == string(jewels[i]) {
				count++
			}
		}
	}

	return count
}

func numJewelsInStonesMap(jewels string, stones string) int {
	dic := make(map[uint8]int)
	count := 0

	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if jewels[i] == stones[j] {
				val, exists := dic[jewels[i]]
				if !exists {
					dic[jewels[i]] = 1
				} else {
					dic[jewels[i]] = val + 1
				}
			}
		}
	}

	for _, val := range dic {
		count += val
	}

	return count
}
