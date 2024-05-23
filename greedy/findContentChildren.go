package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(g) <= 0 || len(s) <= 0 {
		return 0
	}
	sort.Ints(g) // 胃口
	sort.Ints(s) // 饼干

	// 大饼干s尽量满足胃口g的小孩
	sIdx := len(s) - 1
	index := 0
	for gIdx := len(g) - 1; gIdx >= 0; gIdx-- {
		if sIdx >= 0 && s[sIdx] >= g[gIdx] {
			index++
			sIdx--
		}
	}
	return index
}
