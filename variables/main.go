package main 

import (
	"fmt";
)

func main() {
	// boolean
	var a, b, c bool = false, false, true;
	println(a)
	println(b)
	println(c)

	// string
	var string1, string2, string3 string = "go learn 1", "go learn 2", "go learn 3";
	
	println(string1)
	println(string2)
	println(string3)

	// non type
	nonTypeInt, nonTypeString, nonTypeBool := false, 10, "lern go"
	
	println(nonTypeInt)
	println(nonTypeString)
	println(nonTypeBool)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
}


