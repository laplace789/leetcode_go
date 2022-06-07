package backtracking

//Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	backTrack(n, n, "", &res)
	return res

}

//left => nums of `(` , Right => nums of `)`
func backTrack(left int, right int, cur string, res *[]string) {
	// remove `)` before `)` case
	if left > right {
		return
	}

	if left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	//left
	//add
	cur += "("
	//backTrack
	backTrack(left-1, right, cur, res)
	//remove
	cur = cur[:len(cur)-1]

	//right
	//add
	cur += ")"
	//backTrack
	backTrack(left, right-1, cur, res)
	//remove
	cur = cur[:len(cur)-1]
}
