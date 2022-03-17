package main

import "fmt"

func main() {
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 2}))
	fmt.Println(duplicateInArray([]int{1, 4, 7, 2, 3}))
}

func duplicateInArray(arr []int) int {
	visited := make(map[int]bool, 0)

	for _, num := range arr {
		if visited[num] == true {
			fmt.Println(visited)
			return num
		} else {
			visited[num] = true
		}
	}

	return -1
}
