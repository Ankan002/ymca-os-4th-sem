package main

import "fmt"

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < (len(nums) - i); j++ {
			if nums[j-1] > nums[j] {
				var temp int = nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp
			}
		}
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

	bubbleSort(nums)
	
	output_string := fmt.Sprintln("The sorted array is: ", nums)
	fmt.Println(output_string)
}
