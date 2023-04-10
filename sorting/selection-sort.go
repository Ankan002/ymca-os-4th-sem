package main

import "fmt"

func selectionSort(nums []int) {
	for i := 0; i < (len(nums) - 1); i++ {
		currentMinIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[currentMinIndex] {
				currentMinIndex = j
			}
		}

		nums[i], nums[currentMinIndex] = nums[currentMinIndex], nums[i]
	}
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

	selectionSort(nums)
	
	output_string := fmt.Sprintln("The sorted array is: ", nums)
	fmt.Println(output_string)
}
