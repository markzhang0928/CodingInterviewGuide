package main

import "fmt"

func numberToBinary(num int) []int {
	remain := num % 2
	arrayNum := []int{}
	left := num / 2

	if left == 0 {
		return arrayNum
	}
	if remain == 0 {
		arrayNum = append(arrayNum, 0)
		numberToBinary(left)
	} else if remain != 0 {
		arrayNum = append(arrayNum, 1)
		numberToBinary(left)
	}
}

func countNumber(a int) int {
	var count = 0
	if a > 0 {
		result := numberToBinary(a)
		for _, value := range result {
			if value == 1 {
				count++
			}
		}
		return count

	} else if a < 0 {
		result := numberToBinary(a)
		for _, value := range result {
			if value == 0 {
				count++
			}
		}
		return count
	}
	return -1
}
func main() {
	num := countNumber(100)
	fmt.Println("binary count numbers ", num)
}
