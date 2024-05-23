package main

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func printAnimal(word string, ch1 <-chan struct{}, ch2 chan<- struct{}) {
	count := 0
	defer Wg.Done()
	for _ = range ch1 {
		fmt.Println(word)
		count++
		ch2 <- struct{}{}
		if count == 100 {
			fmt.Println(word, "word print end")
			return
		}
	}
}

func main() {
	chCatOk := make(chan struct{}, 1)
	// defer close(chCatOk)
	chDogOk := make(chan struct{})
	// defer close(chDogOk)
	chFishOk := make(chan struct{})
	// defer close(chFishOk)
	Wg.Add(3)
	go printAnimal("cat", chCatOk, chDogOk)
	go printAnimal("dog", chDogOk, chFishOk)
	go printAnimal("fish", chFishOk, chCatOk)
	chCatOk <- struct{}{}
	Wg.Wait()
}
