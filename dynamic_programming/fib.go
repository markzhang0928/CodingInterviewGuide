package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	// 初始化 dp 数组
	prev, curr, sum := 0, 1, 0
	// 遍历顺序: 由前向后。因为后面要用到前面的状态
	for i := 1; i < n; i++ {
		// 确定递归公式/状态转移公式
		sum = prev + curr
		prev, curr = curr, sum
	}
	return sum
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	// dp[i]: 达到i阶有dp[i]种方法
	// 初始化 dp 数组
	dp[1] = 1
	dp[2] = 2
	//  遍历顺序: 由前向后。因为后面要用到前面的状态
	for i := 3; i <= n; i++ {
		// 递推公式
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
