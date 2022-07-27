package sort

//Input: s = "abcd", t = "abcde"
//Output: "e"
//Explanation: 'e' is the letter that was added.

func findTheDifference(s string, t string) byte {

	sByte := []byte(s)
	tByte := []byte(t)

	res := make(map[byte]int)

	for i := 0; i < len(tByte); i++ {
		res[tByte[i]]++
	}

	for i := 0; i < len(sByte); i++ {
		count, exists := res[sByte[i]]
		if exists {
			if count > 1 {
				res[sByte[i]] = count - 1
			} else {
				delete(res, sByte[i])
			}
		}
	}

	//still in map
	for key, _ := range res {
		return key
	}
	return uint8(0)
}
