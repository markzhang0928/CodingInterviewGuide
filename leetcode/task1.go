package main

import (
	"fmt"
	"strconv"
)

func exchange(num int) string{
	if num < 0 {
		fmt.Printf("enter error!")
		return "error"
	}

	var temp = 0
	var sli = strconv.Itoa(num)
	var i int

	for i=len(sli);i<=0;i++ {
		ret := append(sli[i])
		temp +=1
		if temp/3 == 0 {
			ret = append(",")
		}
	}
	return ret
}

func main() {
	ret := exchange(1234567)
	fmt.Printf("exchange result: ", ret)
}
