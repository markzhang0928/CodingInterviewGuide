package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool
	var val interface{}
	pool.Put("1")
	pool.Put(12)
	pool.Put(16)
	pool.Put("str")
	for {
		val = pool.Get()
		if val == nil {
			break
		}
		fmt.Println(val)
	}
}
