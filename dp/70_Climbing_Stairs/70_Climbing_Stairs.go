package dp

/*

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

DP重點 找到那個公式

	f0 = 1
	f1 = 1
	f2 = 1+1
       = 2


	f3 = 1+1+1 = f2+f1 = 3
       =2+1
       =1+2

	f4 = 1+1+1+1 =f3+f2 = 5
       = 2+1+1
       = 1+2+1
       = 1+1+2
       = 2+2

	f5 = 1+1+1+1+1 = f4+f3 = 8
       = 2+1+1+1
       = 1+2+1+1
       = 1+1+2+1
       = 1+1+1+2
       = 2+2+1
       = 1+2+2
       = 2+1+2

*/

func climbStairs(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	arr := make([]int, n+1)

	//先第一第二個
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[len(arr)-1]
}
