package string

/*

Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

Input: s = "Hello"
Output: 1

Input: s = "love live! mu'sic forever"
Output: 4

Input: s = ""
Output: 0

Input: s = "                "
Output: 0

*/

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	nums, isLetter := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if isLetter > 0 {
				nums++
				isLetter = 0
			}
		} else {
			isLetter++
		}
	}
	if isLetter > 0 {
		nums++
	}

	return nums
}
