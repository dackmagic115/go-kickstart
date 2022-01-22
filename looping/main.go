package main

import (
	"fmt"
)

func main() {
	// first()
	// sec()
	// third()
	//arrayLooping()

	stringLoop()
}

func stringLoop() {
	var s string = "Hello, world"

	for index, character := range s {
		fmt.Printf("the character %c in position %d \n", character, index)
	}
}

func first() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func sec() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}

func third() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break
			}
		}
	}
}

func arrayLooping() {
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}
