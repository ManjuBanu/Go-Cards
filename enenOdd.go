package main

import "fmt"

func evenOdd() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(num)

	for _, value := range num {
		if value%2 == 0 {
			fmt.Println(value, "is Even number")
		} else {
			fmt.Println(value, "is Odd number")
		}
	}
}
