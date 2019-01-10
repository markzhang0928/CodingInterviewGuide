package Unit4_recursion_dynamic_programming

/*

给定两个字符串str1和str2，返回两个字符串的最长公共子序列

举例：

str1 = "1A2C3D4B56"

str2 = "B1D23CA45B6A"

它们的最长公共子序列为："123456" 或 "12C4B6"

 */

 // 状态数组 dp
 func GetDp(arr1, arr2 []rune)[][]int {
 	dp := make([][]int, len(arr1))

 	for k, _ := range dp {
 		dp[k] = make([]int, len(arr2))
	}

 	if arr1[0] == arr2[0] {
 		dp[0][0] = 1
	} else {
		dp[0][0] = 0
 	}

 	for i:=1;i<len(arr1);i++{ // 计算第一列
 		if arr1[i] == arr2[0] {
 			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

 	for j:=1;j<len(arr2);j++{ // 计算第一行
		if arr2[j] == arr1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

 	for i:=1;i<len(arr1);i++{
 		for j:=1;j<len(arr2);j++{
 			if arr1[i] == arr2[j] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = getMax(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp
 }


 func getMax(a int, b int) int{
 	if a > b {
 		return a
	} else if a < b {
		return b
	} else {
		return a
	}
 }

func GetRes(dp [][]int, arr1 []rune, arr2 []rune)[]rune{
	m, n := len(arr1)-1, len(arr2)-1
	res := make([]rune, dp[m][n])
	index  := len(res) -1
	for index >= 0 {
		if n>0 && dp[m][n] == dp[m][n-1]{
			n--
		}else if m > 0 && dp[m][n] == dp[m-1][n]{
			m--
		}else{
			res[index] = arr1[m]
			index --
			m--
			n--
		}
	}
	return res
}