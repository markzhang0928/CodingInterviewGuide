package main

// 组合77
var (
	path []int
	res  [][]int
)

func combine(n int, k int) [][]int {
	path, res = make([]int, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, start int) {
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
		if n-i+1 < k-len(path) { // 剪枝 所需需要的元素个数大于剩余的元素个数，直接退出
			break
		}
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}
