package main

import (
	"fmt"
	"math"
)

type IncomingProcess struct {
	ProcessId   int
	ArrivalTime int
	BurstTime   int
}

type ScheduledProcess struct {
	ProcessId      int
	ArrivalTime    int
	BurstTime      int
	CompletionTime int
	WaitingTime    int
	TurnAroundTime int
}

type SPNResponse struct {
	Processes             []ScheduledProcess
	AverageWaitingTime    int
	AverageTurnAroundTime int
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

func getNextProcessIndex(scheduledProcesses []ScheduledProcess, completedSet []bool, currentTime int) int {
	currentSmallestTime := int(math.Inf(+1))
	smallestIndex := -1

	for processIndex, scheduledProcess := range scheduledProcesses {
		if completedSet[processIndex] {
			continue
		}

		if currentTime >= scheduledProcess.ArrivalTime && (smallestIndex == -1 || scheduledProcess.BurstTime < currentSmallestTime) {
			currentSmallestTime = scheduledProcess.BurstTime
			smallestIndex = processIndex
		}
	}

	return smallestIndex
}

func SPNScheduling(incomingProcesses []IncomingProcess) SPNResponse {
	response := SPNResponse{}

	scheduledProcesses := []ScheduledProcess{}

	maxTimeRequired := 0
	current_time_elapsed := 0

	completedSet := []bool{}

	for _, incomingProcess := range incomingProcesses {
		scheduledProcesses = append(scheduledProcesses, ScheduledProcess{
			ProcessId: incomingProcess.ProcessId,
			ArrivalTime: incomingProcess.ArrivalTime,
			BurstTime: incomingProcess.BurstTime,
		})

		maxTimeRequired += incomingProcess.BurstTime
		completedSet = append(completedSet, false)
	}

	for current_time_elapsed < maxTimeRequired {
		nextProcessToBeExecutedIndex := getNextProcessIndex(scheduledProcesses, completedSet, current_time_elapsed)

		if nextProcessToBeExecutedIndex == -1 {
			current_time_elapsed++
			continue
		}

		completedSet[nextProcessToBeExecutedIndex] = true

		scheduledProcesses[nextProcessToBeExecutedIndex].CompletionTime = current_time_elapsed + scheduledProcesses[nextProcessToBeExecutedIndex].BurstTime

		scheduledProcesses[nextProcessToBeExecutedIndex].WaitingTime = current_time_elapsed - scheduledProcesses[nextProcessToBeExecutedIndex].ArrivalTime

		scheduledProcesses[nextProcessToBeExecutedIndex].TurnAroundTime = scheduledProcesses[nextProcessToBeExecutedIndex].WaitingTime + scheduledProcesses[nextProcessToBeExecutedIndex].BurstTime

		current_time_elapsed += scheduledProcesses[nextProcessToBeExecutedIndex].BurstTime
	}

	response.Processes = scheduledProcesses
	response.AverageWaitingTime = getAverageWaitingTime(scheduledProcesses)
	response.AverageTurnAroundTime = getAverageTurnAroundTime(scheduledProcesses)

	return response
}

func main() {
	scheduledProcessesResponse := SPNScheduling([]IncomingProcess{
		{
			ProcessId: 1,
			ArrivalTime: 0,
			BurstTime: 5,
		},
		{
			ProcessId: 2,
			ArrivalTime: 2,
			BurstTime: 4,
		},
		{
			ProcessId: 3,
			ArrivalTime: 3,
			BurstTime: 7,
		},
		{
			ProcessId: 4,
			ArrivalTime: 5,
			BurstTime: 6,
		},
	})

	scheduledProcesses := fmt.Sprintln("Scheduled Processes Are:", scheduledProcessesResponse.Processes)
	averageWaitingTime := fmt.Sprintln("Average Waiting Time is:", scheduledProcessesResponse.AverageWaitingTime)
	averageTurnAroundTime := fmt.Sprintln("Average Turn Around Time is:", scheduledProcessesResponse.AverageTurnAroundTime)

	fmt.Print(scheduledProcesses)
	fmt.Print(averageWaitingTime)
	fmt.Print(averageTurnAroundTime)
}
