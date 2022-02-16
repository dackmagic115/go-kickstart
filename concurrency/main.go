package main

import (
	"fmt"
	"strings"
	"time"
)

// concurrency is code can execute same time
// thread is part of process on cpu
// concurrency on golang use goroutines and channels

func main() {
	fmt.Println("concurrency")
	// sampleGoroutines()
	exampleSendEmail()
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

func exampleSendEmail() {
	// add queue like send email
	// split thread with goroutine for dont want to wait process to send email
	// process send email use so many time
	data := []string{
		"A|20|male",
		"B|20|female",
		"C|20|male",
	}

	for _, person := range data {
		go exampleProcessSendEmail(person)
		fmt.Println("Request has been added to queue.")
	}

	time.Sleep(4 * time.Second)
}

func exampleProcessSendEmail(person string) {
	parts := strings.Split(person, "|")
	name := parts[0]
	age := parts[1]
	gender := parts[2]

	time.Sleep(3 * time.Second) // simulate time consuming task
	fmt.Printf("Send email >> Name=%-5s gender=%-6s age=%2s\n", name, gender, age)
}
