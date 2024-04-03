package main

import (
	"bufio"
	"fmt"
	"os"
)

func stripSentences(str string) (stripedStr string) {
	lengt := len(str)
	if lengt <= 0 {
		return " "
	}
	for i := 0; i < lengt; i++ {
		if i+1 < lengt {
			if str[i] >= 48 && str[i] <= 57 {
				stripedStr += string(str[i])
			} else if str[i] >= 97 && str[i] <= 122 {
				stripedStr += string(str[i])
			} else if str[i] >= 65 && str[i] <= 90 {
				stripedStr += string(str[i])
			} else if str[i] == ' ' || str[i] == '-' {
				stripedStr += " "
			} else if str[i] == '-' && str[i+1] == '-' {
				stripedStr += " "
			}
		}
	}
	return stripedStr
}

func reverseSentences(str string) string {

	lenStr := len(str)
	if lenStr <= 0 {
		return " "
	}
	stripedStr := stripSentences(str)
	return stripedStr
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
		fmt.Println(reverseSentences(string(input)))
	}
}
