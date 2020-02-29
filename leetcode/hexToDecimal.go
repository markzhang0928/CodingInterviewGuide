package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

func latinToDigital(str rune) int {

	if str < 0 {
		return -1
	}

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
	case 'A', 'a':
		return 10
	case 'B', 'b':
		return 11
	case 'C', 'c':
		return 12
	case 'D', 'd':
		return 13
	case 'E', 'e':
		return 14
	case 'F', 'f':
		return 15
	default:
		return -1
	}
}

func hexToDecimal(str string) int {
	if len(str) < 0 {
		return -1
	}

	hexNum := strings.Split(str, "0x")
	ch := []rune(hexNum[1])
	digit := len(ch) - 1
	var sum int
	for _, value := range ch {
		num := latinToDigital(value)
		sum += pow(16, digit) * num
		digit--
	}
	return sum
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := inputReader.ReadLine()
		if err != nil {
			fmt.Errorf("read line error")
		}
		if input == nil {
			break
		}
		fmt.Println(hexToDecimal(string(input)))

	}
}
