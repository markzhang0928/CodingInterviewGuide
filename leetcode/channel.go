package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	chEven := make(chan int, 50)
	chOdd := make(chan int, 50)

	for i := 1; i <= 99; i += 2 {
		// fmt.Println(i)
		chOdd <- i
	}
	for i := 2; i <= 100; i += 2 {
		// fmt.Println(i)
		chEven <- i
	}

	wg.Add(1)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer wg.Done()
		for value := range chOdd {
			ch1 <- value
			fmt.Println("Printer1", value)
			<-ch2
		}
	}()

	for value := range chEven {
		<-ch1
		fmt.Println("Printer2", value)
		ch2 <- value
	}

	wg.Wait()

	close(chOdd)
	close(chEven)
	close(ch1)
	close(ch2)

	time.Sleep(2 * time.Second)
}
