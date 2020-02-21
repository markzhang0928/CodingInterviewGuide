package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	//num := make([]int, 100)
	countCh := make(chan int, 100)
	wg.Add(2)

	go Counter(ch, countCh)
	go Counter(ch, countCh)

	ch <- 1
	wg.Wait()
	go func() {
		time.Sleep(3 * time.Second)
		close(countCh)
	}()
	for value := range countCh {
		fmt.Println("STDOUT:", value)
	}
}

func Counter(ch chan int, count chan int) {
	defer wg.Done()

	for {
		turn, ok := <-ch
		if !ok {
			fmt.Println("channel shut down")
			return
		}

		if turn == 101 {
			fmt.Println("it's the end")
			close(ch)
			return
		}

		if turn%2 == 0 {
			fmt.Println("even number", turn)
		} else if turn%2 == 1 {
			fmt.Println("odd number", turn)
		}

		count <- turn
		turn++
		ch <- turn
	}
}
