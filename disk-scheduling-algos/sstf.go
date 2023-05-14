package main

import (
	"fmt"
	"math"
)

func checkSmallestSeekTimeNextIndex(requestSequence []int, currentHeadPos int, isServiced []bool) int {
	currentSmallestIndex := -1

	for requestIndex, request := range requestSequence {
		if isServiced[requestIndex] {
			continue
		}

		if currentSmallestIndex == -1 || ((math.Abs(float64(request - currentHeadPos))) < math.Abs(float64(requestSequence[currentSmallestIndex]-currentHeadPos))) {
			currentSmallestIndex = requestIndex
		}
	}

	return currentSmallestIndex
}

func sstf(requestSequence []int, initialHeadPos int) {
	serviceSequence := []int{
		initialHeadPos,
	}
	currentHeadPos := initialHeadPos
	totalHeadMovement := 0

	serviceSet := make([]bool, len(requestSequence))

	for serviceNumber := 0; serviceNumber < len(requestSequence); serviceNumber++ {
		nextServiceIndex := checkSmallestSeekTimeNextIndex(requestSequence, currentHeadPos, serviceSet)

		if nextServiceIndex == -1 {
			break
		}

		serviceSequence = append(serviceSequence, requestSequence[nextServiceIndex])

		totalHeadMovement += (int(math.Abs(float64(currentHeadPos - requestSequence[nextServiceIndex]))))

		serviceSet[nextServiceIndex] = true

		currentHeadPos = requestSequence[nextServiceIndex]
	}

	fmt.Println(serviceSequence)
	fmt.Println(totalHeadMovement)
}

func main() {
	sstf([]int{
		176,
		79,
		34,
		60,
		92,
		11,
		41,
		114,
	}, 50)
}
