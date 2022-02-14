package main

import (
	"fmt"
	"time"
)

// concurrency is code can execute same time
// thread is part of process on cpu
// concurrency on golang use goroutines and channels

func main() {
	fmt.Println("concurrency")
	sampleGoroutines() //
}

func sampleGoroutines() {
	fmt.Println("sampleGoroutines")
	// go order follow func gonna create new threads and exactly work
	// run 2 func same time
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Task 2")
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Task 1")
	}

}
