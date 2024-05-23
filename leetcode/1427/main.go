package main

import (
	"strings"
)

func turningRightOrLeft(direction int, amount int, str string) string {
	s := strings.Split(str, "")
	n := len(s)
	if amount == 0 || n == amount || n == 1 {
		return str
	} else if amount > n {
		amount = amount % n
	}
	if direction == 1 {
		//右移
		left := s[n-amount:]
		right := s[0 : n-amount]
		result := append(left, right...)
		return strings.Join(result, "")
	} else {
		right := s[0:amount]
		left := s[amount:]
		result := append(left, right...)
		return strings.Join(result, "")
	}
}

func stringShift(s string, shift [][]int) string {

	for _, obj := range shift {
		direction, amount := obj[0], obj[1]
		s = turningRightOrLeft(direction, amount, s)
	}
	return s
}

func main() {
	s := "abc"
	shift := [][]int{{1, 2}, {0, 3}}
	result := stringShift(s, shift)
	println(result)

}
