package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getLongestDigitalString(str string) string {
	lenStr := len(str)
	if lenStr <= 0 {
		return ""
	}
	// i: leftPtr  j: rightPtr
	// i, j = i+1, j+1
	i, j := 0, 0
	for i < lenStr && j < lenStr {
		if str[j] == ' ' {
			j = j + 1
			i = i + 1
			continue
		}
		//input: abcd123.4567.890.123
		// output: 4567.890
		if str[j] >= 48 && str[j] <= 57 {
			j++
		} else if str[i] == '.' {
			j++
		} else {
			i++
			j++
		}

	}
	subStr := str[i:j]
	maxSubStr := ""
	maxSize := 0
	subStrArr := strings.Split(subStr, ".")
	maxLen := len(subStrArr)
	for i, value := range subStrArr {
		if i+1 < maxLen {
			tmpSize := len(value) + len(subStrArr[i+1])
			if tmpSize >= maxSize {
				maxSubStr = value + "." + subStrArr[i+1]
				maxSize = len(maxSubStr)
			} else {
				continue
			}
		}
	}

	return maxSubStr
}

func main() {
	//input: abcd123.4567.890.123
	// output: 4567.890
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := inputReader.ReadLine()
		if err != nil {
			fmt.Errorf("read line error")
		}
		if input == nil {
			break
		}
		fmt.Println(getLongestDigitalString(string(input)))
	}
}
