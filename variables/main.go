package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// boolean
	var a, b, c bool = false, false, true
	println(a)
	println(b)
	println(c)

	// string
	var string1, string2, string3 string = "go learn 1", "go learn 2", "go learn 3"

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

	// bytes
	myslice := []byte{0x48, 0x65, 0x6C, 0x6C, 0x6f}
	myString := string(myslice)

	fmt.Println(myString)

	// times
	t := time.Now()
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	// map

	mapCreated := make(map[string]int)
	mapCreated["key1"] = 1
	fmt.Println(mapCreated)

	var mapList = map[string]int{"key": 1}
	fmt.Println(mapList)

	var mapAssigned map[string]int
	mapAssigned = mapList
	mapAssigned["key2"] = 2
	fmt.Println(mapAssigned)

	// array & slice
	var stringArray [5]string
	fmt.Println(stringArray)
	fmt.Println(reflect.ValueOf(stringArray).Kind())

	sliceString := []string{"a", "b"}
	sliceString = append(sliceString, "c")
	fmt.Println(sliceString)
	fmt.Println(reflect.ValueOf(sliceString).Kind())

}
