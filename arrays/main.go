package main

import (
	"fmt"
)

func main()  {
	a := []int{}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a,1)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	a = append(a, []int{2, 3, 4, 5}...)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	//array is fixed size 
	//item should max size all time
	aa := []int{1, 2, 3, 4, 5}
	b := aa[:len(aa)-1]
	c := append(aa[:2],aa[3:]...)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(aa[2:3])


}