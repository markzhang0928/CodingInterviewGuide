package main

import "fmt"

func reverseString(s []byte) {
	//left := 0
	max := len(s) - 1
	for left, right := 0, max; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func main() {
	sli := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	reverseString(sli)
	fmt.Println(sli)
}
