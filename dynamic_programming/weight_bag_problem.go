package main

import "fmt"

func test_1_wei_bag_problem1(weight, value []int, bagweight int) int {
	// 定义 and 初始化
	dp := make([]int, bagweight+1)
	// 递推
	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	fmt.Print(dp)
	return dp[bagweight]
}

func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	// 定义dp数组 dp[i][j] [0, i]物品任取放容量为j的背包
	// dp[i][j]
	// 不放物品i , dp[i-1][j]
	// 放物品i , dp[i-1][j-weight[i]] +value[i]

	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test_1_wei_bag_problem1(weight, value, 4)
}
