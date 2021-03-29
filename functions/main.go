package main

import (
	"fmt";
)

func main() {
	// sayMessage("Hello Go!")
	// s := sum(1,2,3,4,5)
	// fmt.Println("The sum is ", s)

	// d, err := divide(5.0, 5.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)


	// g := greeter {
	// 	greeting: "hello",
	// 	name: "go",
	// }
	// g.greet()
	// fmt.Println("The new name is:", g.name)

	func3()
}


func sayMessage(msg string) {
	fmt.Println(msg)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}

	return 
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		panic("Cannot provide zero as second value")
	}

	return a / b, nil
}

type greeter struct {
	greeting string
	name string
}

func (g *greeter) greet(){
	fmt.Println(g.greeting, g.name)
}


func func3() {
	var divide func(float64, float64) (float64, error)
	divide = func(a,b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(d)
}