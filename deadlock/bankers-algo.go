package main

import (
	"fmt"
	"strconv"
)

type BankersAlgorithmResponse struct {
	IsDeadlockPossible bool
	Error              string
	ExecutionSequence  string
}

func checkProcessDone(requirements [][]int) bool {
	for _, requirement := range requirements {
		if requirement[0] != 0 || requirement[1] != 0 || requirement[2] != 0 {
			return false
		}
	}

	return true
}

func bankersAlgorithm(allocated [][]int, max [][]int, available []int) BankersAlgorithmResponse {
	if len(allocated) != len(max) {
		return BankersAlgorithmResponse{
			Error: "Number of processes must be same",
		}
	}

	need := [][]int{}

	for processNumber, maxNeeded := range max {
		need = append(need, []int{
			maxNeeded[0] - allocated[processNumber][0],
			maxNeeded[1] - allocated[processNumber][1],
			maxNeeded[2] - allocated[processNumber][2],
		})
	}

	response := BankersAlgorithmResponse{}

	for true {
		var hasExecutedOne bool = false

		for currentProcessIndex, currentProcessNeed := range need {
			if (currentProcessNeed[0] != 0 || currentProcessNeed[1] != 0 || currentProcessNeed[2] != 0) && currentProcessNeed[0] <= available[0] && currentProcessNeed[1] <= available[1] && currentProcessNeed[2] <= available[2] {
				hasExecutedOne = true

				available[0] += allocated[currentProcessIndex][0]
				currentProcessNeed[0] = 0

				available[1] += allocated[currentProcessIndex][1]
				currentProcessNeed[1] = 0

				available[2] += allocated[currentProcessIndex][2]
				currentProcessNeed[2] = 0

				response.ExecutionSequence += "P" + strconv.Itoa(currentProcessIndex) + " -> "
			}
		}

		if !hasExecutedOne {
			response.IsDeadlockPossible = true

			return response
		}

		if checkProcessDone(need) {
			break
		}
	}

	return response
}

func main() {
	res := bankersAlgorithm([][]int{
		{
			0,
			1,
			0,
		},
		{
			2,
			0,
			0,
		},
		{
			3,
			0,
			2,
		},
		{
			2,
			1,
			1,
		},
		{
			0,
			0,
			2,
		},
	}, [][]int{
		{
			7,
			5,
			3,
		},
		{
			3,
			2,
			2,
		},
		{
			9,
			0,
			2,
		},
		{
			2,
			2,
			2,
		},
		{
			4,
			3,
			3,
		},
	}, []int{
		3,
		3,
		2,
	})

	fmt.Println("****Batch One****")

	if res.Error != "" {
		fmt.Println(res.Error)
	} else if res.IsDeadlockPossible {
		fmt.Println("DEADLOCK");
		fmt.Println("Execution Sequence: " + res.ExecutionSequence + "DEADLOCK")
	} else if !res.IsDeadlockPossible {
		fmt.Println("Execution Sequence: " + res.ExecutionSequence + "END")
	}

	fmt.Println()

	resTwo := bankersAlgorithm([][]int{
		{
			0,
			1,
			0,
		},
		{
			2,
			0,
			0,
		},
		{
			3,
			0,
			2,
		},
		{
			2,
			1,
			1,
		},
		{
			0,
			0,
			2,
		},
	}, [][]int{
		{
			7,
			5,
			3,
		},
		{
			3,
			2,
			2,
		},
		{
			9,
			0,
			2,
		},
		{
			2,
			2,
			2,
		},
		{
			7000,
			3,
			3,
		},
	}, []int{
		3,
		3,
		2,
	})

	fmt.Println("****Batch Two****")

	if resTwo.Error != "" {
		fmt.Println(res.Error)
	} else if resTwo.IsDeadlockPossible {
		fmt.Println("DEADLOCK");
		fmt.Println("Execution Sequence: " + resTwo.ExecutionSequence + "DEADLOCK")
	} else if !resTwo.IsDeadlockPossible {
		fmt.Println("Execution Sequence: " + resTwo.ExecutionSequence + "END")
	}
}
