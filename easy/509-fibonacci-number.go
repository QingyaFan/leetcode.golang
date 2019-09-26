package easy

// fib 返回斐波那契数
// 递归版本
func fib(N int) int {
	if N <= 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}

// fibDP 动态规划的斐波那契
// 执行时间和占用内存都优于递归版本
func fibDP(N int) int {
	var dp []int
	dp = append(dp, 0, 1, 1)
	for i := 3; i < N+1; i++ {
		tmp := dp[i-1] + dp[i-2]
		dp = append(dp, tmp)
	}
	return dp[N]
}
