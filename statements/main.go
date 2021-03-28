package main

import (
	"fmt"
	"math"
)

func main() {
	checkObject()
	statement()
	statementWithMath()
	switchCase()
	swicthCaseWithType() 
}

func checkObject() {
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

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
}

func statement() {
	number := 50 
	guess := -5
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100!")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it!")
		}
	
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	fmt.Println(!false)

}

func statementWithMath() {
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum), 2){
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
}


func switchCase() {
	switch i:= 2+3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default: 
		fmt.Println("another number")
	}
}

func swicthCaseWithType() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is another type")
	}

}