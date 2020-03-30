package main

import (
	"fmt"
)

func power(a, b int) int {
	var result = 1
	for i := 1; i <= b; i++ {
		result *= a
	}
	return result
}

func stringToInt(str rune) int {

	switch str {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		return 0
	}
}

func intToString(a int) string {
	switch a {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	default:
		return "0"
	}
}
func multipleResult(num1, num2 string) string {

	if num1 == "1" {
		return num2
	} else if num1 == "0" || num2 == "0" {
		return "0"
	} else if num2 == "1" {
		return num1
	}
	aStr := []rune(num1)
	bStr := []rune(num2)
	tmp := 0
	tmp2 := 0
	var j = 0
	for _, a := range aStr {
		var i = 0
		for _, b := range bStr {
			tmp += stringToInt(a) * stringToInt(b) * power(10, len(aStr)-i-1)
			i++
			fmt.Println(tmp)
		}
		tmp2 += tmp * power(10, len(bStr)-j-1)
		j++
	}
	len := j
	fmt.Println(tmp2)
	var result string
	for j = len; j >= 0; j-- {
		numStr := tmp2 % 10
		result += intToString(numStr)
	}
	return result
}

func main() {
	var a string = ""
	var b string = ""
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%s\n", multipleResult(a, b))
		}
	}
}
