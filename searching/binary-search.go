package main

import (
	"fmt"
	"strconv"
)

func binarySearch(nums []int, target int) int {
	var startIndex int = 0
	var endIndex int = len(nums) - 1

	for startIndex <= endIndex {
		var midIndex int = startIndex + int((endIndex-startIndex)/2)

		if nums[midIndex] == target {
			return midIndex
		} else if nums[midIndex] > target {
			endIndex = midIndex - 1
		} else {
			startIndex = midIndex + 1
		}
	}

	return -1
}

func main() {
	var nums []int = []int {
		3,
		5,
		8,
		10,
		23,
		24,
		79,
	}

	fmt.Println("The given number is at index: " + strconv.Itoa(binarySearch(nums, 8)))
	fmt.Println("The given number is at index: " + strconv.Itoa(binarySearch(nums, 100)))
}
