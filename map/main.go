package main

import (
	"fmt"
)

func main(){
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illionois": 12801539,
		"Ohio": 11614373,
	}

	fmt.Println(statePopulations)
	// deleted
	delete(statePopulations, "Ohio")
	fmt.Println(statePopulations)

	// search 
	pop, ok := statePopulations["Texas"]
	fmt.Println(pop, ok)


	// len
	fmt.Println(len(statePopulations))
}