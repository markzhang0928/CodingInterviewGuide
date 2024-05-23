package main

func combinationSum3(k int, n int) [][]int {
	path, res = make([]int, 0, k), make([][]int, 0)
	backtracing(k, n, 1, 0)
	return res
}

func backtracing(k int, n int, start int, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	for i := start; i <= 9; i++ {
		if sum+i > n || 9-i+1 < k-len(path) { // 所需要的元素大于 剩余的元素
			break
		}
		path = append(path, i)
		backtracing(k, n, i+1, sum+i)
		path = path[:len(path)-1]
	}
}
