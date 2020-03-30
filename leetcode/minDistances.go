package main

import "math"

func minDynDistance(arr []int, num1, num2 int) int {
	if arr == nil || len(arr) <= 0 {
		return math.MaxInt64
	}
	lastPos1 := -1
	lastPos2 := -1
	minDis := math.MaxInt64
	for i, value := range arr {
		if value == num1 {
			lastPos1 = i
			if lastPos2 >= 0 {
				minDis = Min(minDis, lastPos1-lastPos2)
			}
		}
		if value == num2 {
			lastPos2 = i
			if lastPos1 >= 0 {
				minDis = Min(minDis, lastPos2-lastPos1)
			}
		}
	}
	return minDis
}

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

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
