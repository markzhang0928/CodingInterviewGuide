package main

// pdd second
import "fmt"

func swapSort(arr [7]int) {
	length := len(arr)
	for i := 0; i < length; {
		tmp := arr[i]
		arr[i], arr[tmp-1] = arr[tmp-1], arr[i]
		if arr[i] == i+1 {
			i++
		}
	}
	fmt.Print(arr)
}

func main() {

	arr := [7]int{7, 5, 3, 6, 4, 2, 1}

	swapSort(arr)
}
