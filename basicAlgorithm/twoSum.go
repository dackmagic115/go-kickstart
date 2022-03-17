package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if v, found := m[target-num]; found {
			return []int{nums[v], num}
		}

		m[num] = idx
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{nums[i], nums[j]}
			}
		}
	}

	return []int{}
}
