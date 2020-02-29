package main

import "math"

func minDistance(arr []int, num1, num2 int) int {

	if arr == nil || len(arr) <= 0 {
		return math.MaxInt64
	}

	minDis := math.MaxInt64
	dist := 0
	for i, v := range arr {
		if v == num1 {
			for j, v := range arr {
				if v == num2 {
					dist = Abs(i - j)
				}
				if dist < minDis {
					minDis = dist
				}
			}
		}
	}
	return minDis
}

func Abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}
