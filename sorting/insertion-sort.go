package main

import "fmt"

func insertionSort(nums []int) {
	for i:=1; i<len(nums); i++ {
		current_num := nums[i]

		var current_search_index int = i
		for current_search_index > 0 && nums[current_search_index - 1] >= current_num {
			nums[current_search_index] = nums[current_search_index - 1]
			current_search_index--
		}

		nums[current_search_index] = current_num
	}
}

func main()  {
	var nums []int = []int{
		23,
		45,
		11,
		109,
		105,
		56,
		43,
	}

	insertionSort(nums)
	
	output_string := fmt.Sprintln("The sorted array is: ", nums)
	fmt.Println(output_string)
}