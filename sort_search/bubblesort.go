package main

import "fmt"

// 稳定 O(n2)
func bubbleSort(sli []int) []int {
	len := len(sli)

	if len <= 1 {
		return sli
	}

	//该层循环控制 需要冒泡的轮数
	for i := 1; i < len; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < len-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	return sli
}

func bubbleSort2(sli []int) []int {
	len := len(sli)

	if len <= 1 {
		return sli
	}

	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			if sli[j] < sli[j-1] {
				sli[j], sli[j-1] = sli[j-1], sli[j]
			}
		}
	}
	return sli
}

func main() {
	var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	res := bubbleSort2(sli)
	fmt.Println(res)
}
