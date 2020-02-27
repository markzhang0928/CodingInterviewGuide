package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countSubstr(str string, substr string) int {

	// 所有字符转换成小写
	str = strings.ToLower(str)
	substr = strings.ToLower(substr)
	var count = 0
	runeString := []rune(str)
	for i, _ := range runeString {
		if string(runeString[i]) == substr {
			count++
		}
	}
	return count
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	inputString, _, err := inputReader.ReadLine()
	if err != nil {
		fmt.Errorf("read line error")
	}
	substr, _, err := inputReader.ReadLine()
	if err != nil {
		fmt.Errorf("read line error")
	}
	result := countSubstr(string(inputString), string(substr))
	fmt.Print(result)
}
