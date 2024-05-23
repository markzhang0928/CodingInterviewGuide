package main

func main() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 12
	}()
	// ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	n := <-ch1
	println(n)
}
