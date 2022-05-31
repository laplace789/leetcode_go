package stack

/*

Input: s = "ab#c", t = "ad#c"
Output: true
Explanation: Both s and t become "ac".

Input: s = "ab##", t = "c#d#"
Output: true
Explanation: Both s and t become "".

"xywrrmp"
"xywrrmu#p"

Input: s = "xywrrmp", t = "xywrrmu#p"
Output: true
Explanation: Both s and t become "".

*/

//backspaceCompare very bad solution
func backspaceCompare(s string, t string) bool {
	charS := []rune(s)
	charT := []rune(t)

	stackS := make([]rune, 0)
	stackT := make([]rune, 0)
	for i := 0; i < len(charS); i++ {
		if charS[i] == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, charS[i])
		}
	}

	for i := 0; i < len(charT); i++ {
		if charT[i] == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, charT[i])
		}
	}

	return string(stackT) == string(stackS)
}

func backspaceCompareBetterWay(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	is, js := 0, 0
	for i >= 0 || j >= 0 {
		//find next character in s
		for i >= 0 {
			if s[i] == '#' {
				//move to next index and add a skip
				i--
				is++
			} else if is > 0 {
				i--
				is--
			} else {
				break
			}
		}

		//find next character in t
		for j >= 0 {
			if t[j] == '#' {
				//move to next index and add a skip
				j--
				js++
			} else if js > 0 {
				j--
				js--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		if (i >= 0) != (j >= 0) {
			return false
		}
		i--
		j--
	}

	return true
}
