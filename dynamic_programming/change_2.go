package main

func change(amount int, coins []int) int {

	// 定义dp数组, dp[j] 装背包j, dp[j]种方法
	dp := make([]int, amount+1)
	// 初始化,0大小的背包, 当然是不装任何东西了, 就是1种方法
	dp[0] = 1
	// 遍历顺序
	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			// 推导公式
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
