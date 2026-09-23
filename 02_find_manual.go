package main

import "fmt"

func main() {
	numbers := []int{12, 7, 25, 18, 9}
	target := 100

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Checking index", i, "value", numbers[i])

		if numbers[i] == target {
			fmt.Println("Target found at index", i)
		}
	}
}
