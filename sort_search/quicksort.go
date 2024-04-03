package main

import "fmt"

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

/*
 快排和冒泡 都是每一趟排序把一个元素放置到其最终位置上
*/

// 不稳定 O(nlgN)
func quickSort(sli []int) []int {

	sliLen := len(sli)
	if sliLen <= 1 {
		return sli
	}

	// 选择第一个元素作为基准
	baseNum := sli[0]
	// 遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	// 初始化左右两个切片
	leftSli := []int{}  //小于基准的
	rightSli := []int{} //大于基准的

	for i := 1; i < sliLen; i++ {
		if baseNum > sli[i] {
			leftSli = append(leftSli, sli[i])
		} else {
			rightSli = append(rightSli, sli[i])
		}
	}

	leftSli = quickSort(leftSli)
	rightSli = quickSort(rightSli)

	leftSli = append(leftSli, baseNum)
	return append(leftSli, rightSli...)
}

func main() {
	res := quickSort(sli)
	fmt.Println(res)
}
