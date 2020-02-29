package main

import "fmt"

type StringLi struct {
	startIndex int
	len        int
}

func (p *StringLi) expandBothSide(str string, c1, c2 int) {

	n := len(str)
	for c1 >= 0 && c2 < n && str[c1] == str[c2] {
		c1--
		c2++
	}
	tmpStartIndex := c1 + 1
	tmpLen := c2 - c1 - 1
	if tmpLen > p.len {
		p.len = tmpLen
		p.startIndex = tmpStartIndex
	}
}

func (p *StringLi) getLongestPalidrome(str string) {

	n := len(str)
	if n < 1 {
		return
	}

	for i, _ := range str {
		p.expandBothSide(str, i, i)
		p.expandBothSide(str, i, i+1)
	}
}

func main() {
	str := "abcdefgfedxyz"
	t := &StringLi{}
	t.getLongestPalidrome(str)
	if t.startIndex != -1 && t.len != -1 {
		fmt.Println("the longest palidrome: ", string(str[t.startIndex:t.startIndex+t.len]))
	}
}
