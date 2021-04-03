package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	goruntines()
	goruntines2()
}

func goruntines() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "GoodBye"
	time.Sleep(100 * time.Millisecond)
}

// runtines group
var wg = sync.WaitGroup{}

func goruntines2() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}