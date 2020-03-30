package solution

func partition(arr []int, low, high int) int {
	key := arr[low]
	for low < high {
		for low < high && arr[high] > key {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] < key {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = key
	return low
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	lenNum1 := len(nums1)
	lenNum2 := len(nums2)

	arr := make([]int, lenNum1+lenNum2)
	if lenNum1 == 0 && lenNum2 == 0 {
		return -1
	} else if lenNum1 == 0 && lenNum2 != 0 {
		arr = append(arr, nums2...)
	} else if lenNum2 == 0 && lenNum1 != 0 {
		arr = append(arr, nums1...)
	}

	low := 0
	n := len(arr)
	high := n - 1
	mid := (low + high) / 2

	for true {
		pos := partition(arr, low, high)
		if pos == mid {
			break
		} else if pos > mid {
			high = pos - 1
		} else {
			low = pos + 1
		}
	}
	if (n % 2) != 0 {
		return float64(arr[mid])
	} else {
		return float64((arr[mid] + arr[mid+1]) / 2)
	}
}
