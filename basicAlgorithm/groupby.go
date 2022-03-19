package main

import "fmt"

type user struct {
	id        int
	username  string
	firstname string
}

func main() {
	users := &[]user{{
		id:        1,
		username:  "a",
		firstname: "aa",
	},
		{
			id:        2,
			username:  "b",
			firstname: "aa",
		}, {
			id:        3,
			username:  "c",
			firstname: "cc",
		},
		{
			id:        4,
			username:  "d",
			firstname: "dd",
		},
	}

	groupbyUsername(*users)
	groupByCount(*users)
}

func groupbyUsername(users []user) {
	group := make(map[string][]user)

	for _, user := range users {
		if v, exists := group[user.firstname]; exists {
			v = append(v, user)
		}
		group[user.firstname] = append(group[user.firstname], user)
	}

	fmt.Println(group)
}

func groupByCount(users []user) {
	group := make(map[string]int)

	for _, user := range users {
		group[user.firstname] += 1
	}
	fmt.Println(group)
}
