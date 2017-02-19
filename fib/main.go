package main

import (
	"sync"
	"fmt"
	"time"

	"runtime"
)






func main() {
	runtime.GOMAXPROCS(8)
	n := 50
	start := time.Now()
	fmt.Println(gofib(n))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(fib(n))
	fmt.Println(time.Since(start))
}


func fib(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func gofib(x int) int {
	var wg sync.WaitGroup
	var a, b int
	wg.Add(2)
	if x == 0 || x == 1 {
		return 1
	}
	go func(){
		a = fib(x-1)
		wg.Done()
	}()

	go func(){
		b = fib(x-2)
		wg.Done()
	}()
	wg.Wait()
	return a + b
}