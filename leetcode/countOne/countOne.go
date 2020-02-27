package main

import "fmt"

func countOne(num int) int {

	count := 0
	for num > 0 {
		if (num & 1) == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func main() {
	num := countOne(8)
	fmt.Println("binary count numbers ", num)
}
