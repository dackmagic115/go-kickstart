package main

import (
	"fmt"
	"sync"
	"time"
)



var wg = sync.WaitGroup{}

func main() {
	channel()
	logging()
} 

func channel() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <- ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()
	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct  {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50) // create channel

func logging() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}


func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}