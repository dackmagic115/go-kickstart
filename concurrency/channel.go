package main

import (
	"fmt"
	"time"
)

// go channel concept is blocking
// <- c read from channel
// waiting for write channel
// if read only none data on channel program we call dead lock mode
// thread all sleep
// c <- 1 write to channel
func main() {
	fmt.Println("with channel")

	// channel1()
	channelCaseBlocking()
}

func channel1() {
	channel := make(chan int)
	a := 10
	b := 5

	go add(channel, a, b)
	go multiply(channel)

	time.Sleep(1 * time.Second)
}

func add(c chan int, a int, b int) {
	result := a + b
	c <- result // write channel
}

func multiply(c chan int) {
	result := <-c // read channel
	result *= 2

	fmt.Printf("Result is %d\n", result)
}

func channelCaseBlocking() {
	c := make(chan bool)

	go routine1(c)
	go routine2(c)

	time.Sleep(3 * time.Second)
}

func routine1(c chan bool) {
	fmt.Println("#1 has started, waiting for #2 to start")
	<-c // read channel is blocking before write channel
	fmt.Println("#1 received a notification from #2")
}

func routine2(c chan bool) {
	fmt.Println("#2 has started, do some work and notify #1")
	time.Sleep(2 * time.Second) // simulate working
	c <- true                   // write channel
	fmt.Println("#2 has finished")
}

func channelDeadLockCase() {
	c := make(chan bool)
	// deadlock, channel have not been write
	<-c // read channel

	fmt.Println("DONE") // never reach this print
}

// if create goroutine more
// cpu used is overhead

func channelCreatelimitSlot() {
	c := make(chan bool, 3)
	// if thread more than 4 thread is block
	for true {
		go func() {
			fmt.Println(time.Now().Second())
			time.Sleep(2 * time.Second)
			<-c
		}()

		c <- true
	}
}
