package main

import (
	"fmt"
	"sync"
)

func main() {

	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <- number:
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
}
