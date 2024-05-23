package main

import (
	"fmt"
	"strconv"
	"strings"
)

func switchString(before string) string {
	switch before {
	case "0":
		return "0"
	case "1":
		return "1"
	case "6":
		return "9"
	case "8":
		return "8"
	case "9":
		return "6"
	}
	return before
}

func confusingNumber(n int) bool {
	confusingList := "23457"
	str := strconv.Itoa(n)
	var reverseStr string
	for _, i := range str {
		if strings.ContainsRune(confusingList, rune(i)) {
			return false
		} else {
			reverseStr += switchString(string(i))
		}
		if str == reverseStr {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(confusingNumber(11))
}
