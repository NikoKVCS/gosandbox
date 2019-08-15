package main

import (
	"fmt"
	"gpool"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	pool := gpool.New(100)
	println(runtime.NumGoroutine())
	for i := 0; i < 1000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println(runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	println(runtime.NumGoroutine())
}

var ch chan int
var wg = &sync.WaitGroup{}

func print() {
	fmt.Println(<-ch)
	wg.Done()
}

func main() {
	ch = make(chan int, 10)
	for i := 0; i < 1000; i++ {
		ch <- i
		wg.Add(1)
		go print()
	}
	wg.Wait()
}
