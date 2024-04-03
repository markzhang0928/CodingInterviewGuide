package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func reverse(str string) string {

	if len(str) == 0 {
		return ""
	}
	ch := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		ch[i], ch[j] = ch[j], ch[i]
	}
	return string(ch)
}

func reverseSentences(str string) string {

	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return str
	}
	var reverseStr, reversedStr string
	reverseStr = reverse(str)
	reg := regexp.MustCompile(`\s+`)
	subStrSlice := reg.Split(reverseStr, -1)
	// subStrSlice := strings.Split(reverseStr, " ")
	for i, _ := range subStrSlice {
		reversedStr += reverse(subStrSlice[i]) + " "
	}
	return reversedStr
}

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	input, _, err := inputReader.ReadLine()
	if err != nil {
		fmt.Errorf("read line error")
	}
	fmt.Printf(reverseSentences(string(input)))
}
