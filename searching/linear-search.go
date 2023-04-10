package main

import (
	"fmt"
	"strconv"
)

func linearSearch(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}

	return -1
}

func main() {
	var nums []int = []int{
		23,
		45,
		11,
		109,
		105,
		56,
		43,
	}

	fmt.Println("The given number is at index: " + strconv.Itoa(linearSearch(nums, 1000)))
	fmt.Println("The given number is at index: " + strconv.Itoa(linearSearch(nums, 105)))
}
