package main

import "fmt"

func linearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}

	return -1
}

func main() {
	numbers := []int{12, 7, 25, 18, 9}

	result := linearSearch(numbers, 50)
	fmt.Println("Result index:", result)
}
