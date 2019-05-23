package main

import "fmt"

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func insertSort(sli []int) []int {
	len := len(sli)

	if len <= 1 {
		return sli
	}

	for i := 0; i < len; i++ {
		tmp := sli[i]
		// 内层循环控制,比较并插入
		for j := i - 1; j >= 0; j-- {
			fmt.Println("sli: i=", i, "j= ", j, sli)
			if tmp < sli[j] {
				// 发现插入的元素要小, 交换位置, 将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				// 如果碰到不需要移动的元素,则前面的就不需要再次比较了
				break
			}
		}
	}
	return sli
}

func main() {
	res := insertSort(sli)
	fmt.Println(res)
}
