package main

import "fmt"

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
	TurnAroundTime int
	WaitingTime    int
}

type FCFSResponse struct {
	ScheduledProcesses    []ScheduledProcess
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

func fcfs(incomingProcesses []IncomingProcess) FCFSResponse {
	var response FCFSResponse = FCFSResponse{}
	var calculatedProcesses []ScheduledProcess = []ScheduledProcess{}
	var lastProcessCompletionTime int = 0

	for _, incomingProcess := range incomingProcesses {
		calculatedProcesses = append(calculatedProcesses, ScheduledProcess{
			ProcessId:      incomingProcess.ProcessId,
			ArrivalTime:    incomingProcess.ArrivalTime,
			BurstTime:      incomingProcess.BurstTime,
			CompletionTime: lastProcessCompletionTime + incomingProcess.BurstTime,
			WaitingTime:    lastProcessCompletionTime - incomingProcess.ArrivalTime,
			TurnAroundTime: (lastProcessCompletionTime - incomingProcess.ArrivalTime) + incomingProcess.BurstTime,
		})

		lastProcessCompletionTime += incomingProcess.BurstTime
	}

	response.ScheduledProcesses = calculatedProcesses
	response.AverageWaitingTime = getAverageWaitingTime(calculatedProcesses)
	response.AverageTurnAroundTime = getAverageTurnAroundTime(calculatedProcesses)

	return response
}

func main() {
	var incomingProcesses []IncomingProcess = []IncomingProcess{
		{
			ProcessId:   1,
			ArrivalTime: 0,
			BurstTime:   5,
		},
		{
			ProcessId:   2,
			ArrivalTime: 2,
			BurstTime:   4,
		},
		{
			ProcessId:   3,
			ArrivalTime: 3,
			BurstTime:   7,
		},
		{
			ProcessId:   4,
			ArrivalTime: 5,
			BurstTime:   6,
		},
	}

	scheduledProcessesResponse := fcfs(incomingProcesses)

	scheduledProcesses := fmt.Sprintln("Scheduled Processes Are:", scheduledProcessesResponse.ScheduledProcesses)
	averageWaitingTime := fmt.Sprintln("Average Waiting Time is:", scheduledProcessesResponse.AverageWaitingTime)
	averageTurnAroundTime := fmt.Sprintln("Average Turn Around Time is:", scheduledProcessesResponse.AverageTurnAroundTime)

	fmt.Print(scheduledProcesses)
	fmt.Print(averageWaitingTime)
	fmt.Print(averageTurnAroundTime)
}
