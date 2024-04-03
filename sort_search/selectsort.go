package main

import "fmt"

func selectSort(sli []int) []int {

	length := len(sli)

	if length <= 1 {
		return sli
	}

	for i := 0; i < length-1; i++ {
		flag := i
		for j := i + 1; j < length; j++ {
			if sli[j] < sli[flag] {
				flag = j
			}
		}
		if flag != i {
			sli[i], sli[flag] = sli[flag], sli[i]
		}
	}
	return sli
}

func main() {
	var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	res := selectSort(sli)
	fmt.Println("+++", res)
}
