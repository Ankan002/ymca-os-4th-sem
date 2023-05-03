package main

import "fmt"

type IncomingProcess struct {
	ProcessId      int
	ArrivalTime    int
	BurstTime      int
	PriorityNumber int
}

type ScheduledProcess struct {
	ProcessId      int
	ArrivalTime    int
	BurstTime      int
	PriorityNumber int
	CompletionTime int
	TurnAroundTime int
	WaitingTime    int
}

type PriorityResponse struct {
	Processes             []ScheduledProcess
	AverageWaitingTime    int
	AverageTurnAroundTime int
}

func nextProcessToBeExecutedIndex(scheduledProcesses []ScheduledProcess, currentTime int, remainingTimeForEachProcesses []int) int {
	var highestPriorityIndex int = 0

	for currentIndex, scheduledProcess := range scheduledProcesses {
		if remainingTimeForEachProcesses[highestPriorityIndex] <= 0 {
			highestPriorityIndex = currentIndex
		} else if remainingTimeForEachProcesses[currentIndex] > 0 && scheduledProcess.ArrivalTime <= currentTime && scheduledProcess.PriorityNumber < scheduledProcesses[highestPriorityIndex].PriorityNumber {
			highestPriorityIndex = currentIndex
		}
	}

	return highestPriorityIndex
}

func getAverageWaitingTime(scheduledProcesses []ScheduledProcess) int {
	var totalWaitingTime int = 0

	for _, scheduledProcess := range scheduledProcesses {
		totalWaitingTime += scheduledProcess.WaitingTime
	}

	return totalWaitingTime / len(scheduledProcesses)
}

func getAverageTurnAroundTime(scheduledProcesses []ScheduledProcess) int {
	var totalTurnAroundTime int = 0

	for _, scheduledProcess := range scheduledProcesses {
		totalTurnAroundTime += scheduledProcess.TurnAroundTime
	}

	return totalTurnAroundTime / len(scheduledProcesses)
}

// ! NOTE: This algorithm is only valid if the the starting time of a process is zero.
func priorityQueueScheduling(incomingProcesses []IncomingProcess) PriorityResponse {
	response := PriorityResponse{}

	var totalTimeCyclesNeeded int = 0
	var remainingTimeForEachProcesses []int = []int{}
	var scheduledProcesses []ScheduledProcess = []ScheduledProcess{}

	for _, incomingProcess := range incomingProcesses {
		totalTimeCyclesNeeded += incomingProcess.BurstTime

		remainingTimeForEachProcesses = append(remainingTimeForEachProcesses, incomingProcess.BurstTime)
		scheduledProcesses = append(scheduledProcesses, ScheduledProcess{
			ProcessId:      incomingProcess.ProcessId,
			ArrivalTime:    incomingProcess.ArrivalTime,
			BurstTime:      incomingProcess.BurstTime,
			PriorityNumber: incomingProcess.PriorityNumber,
		})
	}

	var lastProcessExecutedIndex int

	for i := 0; i < totalTimeCyclesNeeded; i++ {
		var currentProcessIndex int

		if i != 0 && scheduledProcesses[lastProcessExecutedIndex].PriorityNumber == 1 && remainingTimeForEachProcesses[lastProcessExecutedIndex] > 0 {
			currentProcessIndex = lastProcessExecutedIndex
		} else {
			currentProcessIndex = nextProcessToBeExecutedIndex(scheduledProcesses, i, remainingTimeForEachProcesses)
			lastProcessExecutedIndex = currentProcessIndex
		}

		for scheduledProcessIndex, scheduledProcess := range scheduledProcesses {
			if scheduledProcessIndex != currentProcessIndex && scheduledProcess.ArrivalTime <= i && remainingTimeForEachProcesses[scheduledProcessIndex] > 0 {
				scheduledProcesses[scheduledProcessIndex].WaitingTime += 1
			}

			if remainingTimeForEachProcesses[scheduledProcessIndex] > 0 {
				scheduledProcesses[scheduledProcessIndex].CompletionTime += 1
			}

		}

		remainingTimeForEachProcesses[currentProcessIndex] -= 1
	}

	for scheduledProcessIndex := range scheduledProcesses {
		scheduledProcesses[scheduledProcessIndex].TurnAroundTime = scheduledProcesses[scheduledProcessIndex].BurstTime + scheduledProcesses[scheduledProcessIndex].WaitingTime
	}

	response.Processes = scheduledProcesses
	response.AverageTurnAroundTime = getAverageTurnAroundTime(scheduledProcesses)
	response.AverageWaitingTime = getAverageWaitingTime(scheduledProcesses)

	return response
}

func main() {
	incomingProcesses := []IncomingProcess{
		{
			ProcessId:      1,
			ArrivalTime:    0,
			BurstTime:      11,
			PriorityNumber: 3,
		},
		{
			ProcessId:      2,
			ArrivalTime:    5,
			BurstTime:      28,
			PriorityNumber: 1,
		},
		{
			ProcessId:      3,
			ArrivalTime:    12,
			BurstTime:      2,
			PriorityNumber: 4,
		},
		{
			ProcessId:      4,
			ArrivalTime:    2,
			BurstTime:      10,
			PriorityNumber: 2,
		},
		{
			ProcessId:      5,
			ArrivalTime:    9,
			BurstTime:      16,
			PriorityNumber: 5,
		},
	}

	scheduledProcessesResponse := priorityQueueScheduling(incomingProcesses)

	scheduledProcesses := fmt.Sprintln("Scheduled Processes Are:", scheduledProcessesResponse.Processes)
	averageWaitingTime := fmt.Sprintln("Average Waiting Time is:", scheduledProcessesResponse.AverageWaitingTime)
	averageTurnAroundTime := fmt.Sprintln("Average Turn Around Time is:", scheduledProcessesResponse.AverageTurnAroundTime)

	fmt.Print(scheduledProcesses)
	fmt.Print(averageWaitingTime)
	fmt.Print(averageTurnAroundTime)
}
