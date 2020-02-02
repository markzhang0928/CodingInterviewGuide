package main

import "fmt"

func merge(sli []int, low, mid, high int) {
	// sli[low.. mid] 和 sli[mid.. high] 各自有序
	lower := mid - low + 1
	higher := high - mid
	L := make([]int, lower)
	for i, k := 0, low; i < lower; i, k = i+1, k+1 {
		L[i] = sli[k]
	}
	H := make([]int, higher)
	for i, k := 0, mid+1; i < higher; i, k = i+1, k+1 {
		H[i] = sli[k]
	}

	var i, k, j int
	for i, k, j = 0, low, 0; i < lower && j < higher; k++ {
		if L[i] < H[j] { // 比较sli 的左右两段中的元素
			sli[k] = L[i]
			i++
		} else {
			sli[k] = H[j]
			j++
		}
	} // for
	if i < lower {
		for j = i; j < lower; j, k = j+1, k+1 {
			sli[k] = L[j]
		}
	}
	if j < higher {
		for i = j; i < higher; i, k = i+1, k+1 {
			sli[k] = H[i]
		}
	}
}

func mergeSort(sli []int, low, high int) {

	if low < high {
		mid := (low + high) / 2
		mergeSort(sli, low, mid)
		mergeSort(sli, mid+1, high)
		merge(sli, low, mid, high)
	}
}

func main() {
	var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	mergeSort(sli, 0, len(sli)-1)
	fmt.Println(sli)
}
