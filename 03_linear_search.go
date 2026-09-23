package main

import "fmt"

func main() {
	numbers := []int{12, 7, 25, 18, 9}
	target := 18

	foundIndex := -1

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			foundIndex = i
			break
		}
	}

	if foundIndex != -1 {
		fmt.Println("Found at index:", foundIndex)
	} else {
		fmt.Println("Target not found")
	}
}
