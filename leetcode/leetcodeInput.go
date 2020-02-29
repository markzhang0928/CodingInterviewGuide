package main

import (
	"bufio"
	"fmt"
	"os"
)

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

//func main() {
//
//	inputReader := bufio.NewReader(os.Stdin)
//	input, _, err := inputReader.ReadLine()
//	if err != nil {
//		fmt.Errorf("read line error")
//	}
//	fmt.Print(countLastWord(string(input)))
//}
