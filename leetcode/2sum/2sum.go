package main

import "fmt"

func TwoSum1(nums []int, target int) []int {
	n := len(nums)
	for i, value := range nums {
		for j := i + 1; j < n; j++ {
			if value+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSum2(nums []int, target int) []int {

	m := make(map[int]int, len(nums))
	for i, value := range nums {
		sub := target - value
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[value] = i
		}
	}
	return nil
}

func main() {
	r1 := TwoSum1([]int{2, 7, 11, 15}, 9)
	fmt.Println(r1)
	r2 := TwoSum2([]int{1, 7, 11, 8}, 9)
	fmt.Println(r2)
}
