package main

import "fmt"

func linearSearch(numbers []int, target int) (int, int) {
	comparisons := 0

	for i := 0; i < len(numbers); i++ {
		comparisons++

		if numbers[i] == target {
			return i, comparisons
		}
	}

	return -1, comparisons
}

func binarySearch(numbers []int, target int) (int, int) {
	left := 0
	right := len(numbers) - 1
	comparisons := 0

	for left <= right {
		middle := (left + right) / 2
		comparisons++

		if numbers[middle] == target {
			return middle, comparisons
		} else if numbers[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1, comparisons
}

func main() {
	studentIDs := []int{101, 105, 108, 112, 119, 125, 131, 140, 155}

	target := 125

	linearIndex, linearComparisons := linearSearch(studentIDs, target)
	binaryIndex, binaryComparisons := binarySearch(studentIDs, target)

	fmt.Println("Target:", target)
	fmt.Println("Linear Search - Index:", linearIndex)
	fmt.Println("Linear Search - Comparisons:", linearComparisons)
	fmt.Println("Binary Search - Index:", binaryIndex)
	fmt.Println("Binary Search - Comparisons:", binaryComparisons)

	fmt.Println()

	target = 130

	linearIndex, linearComparisons = linearSearch(studentIDs, target)
	binaryIndex, binaryComparisons = binarySearch(studentIDs, target)

	fmt.Println("Target:", target)
	fmt.Println("Linear Search - Index:", linearIndex)
	fmt.Println("Linear Search - Comparisons:", linearComparisons)
	fmt.Println("Binary Search - Index:", binaryIndex)
	fmt.Println("Binary Search - Comparisons:", binaryComparisons)
}
