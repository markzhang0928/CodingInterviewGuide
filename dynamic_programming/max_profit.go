package main

// 贪心
func maxProfit(prices []int) int {
	lowest := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-lowest > res {
			res = prices[i] - lowest
		}
		if lowest > prices[i] {
			lowest = prices[i]
		}
	}
	return res
}

// 动规
func maxProfitDP1(prices []int) int {

	length := len(prices)
	if length == 0 {
		return 0
	}
	// dp[i][0] 表示第i天持有股票所得最多现金
	// dp[i][1] 表示第i天不持有股票所得最多现金
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	// 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	// 推导公式：dp[i][0] = max(dp[i - 1][0], -prices[i])
	// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0]+prices[i])
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

// 动规: 滚动数组
func maxProfitDP(prices []int) int {
	// dp[i][0] 表示第i天持有股票所得最多现金
	// dp[i][1] 表示第i天不持有股票所得最多现金
	dp := [2][2]int{}
	// 初始化
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		// 推导公式：dp[i][0] = max(dp[i - 1][0], -prices[i])
		// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0]+prices[i])
		dp[i%2][0] = max(dp[(i-1)%2][0], -prices[i])
		dp[i%2][1] = max(dp[(i-1)%2][1], dp[(i-1)%2][0]+prices[i])
	}

	return dp[(len(prices)-1)%2][1]
}
