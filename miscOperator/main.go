package main

import "fmt"

// pointers
func main() {
	a := 4

	fmt.Println(a)
	changeValue(&a) // pass address of a &a is hex
	fmt.Println(a)

}

func changeValue(value *int) {
	*value = 10 // *value is value of &a
}
