package main

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/printers/taskprinter"
)

// StepsSequencePerformer contains array of AlgorithmStep handlers
// in the order of their start
type StepsSequencePerformer struct {
	steps *[]AlgorithmStep
	task  *taskmodel.Task
}

// Run starts all the AlgorithmStep handlers
// TODO: Refactor printing add logging at some separate object
func (ssp *StepsSequencePerformer) Run() (err error) {
	for i, step := range *ssp.steps {
		fmt.Printf("\n=== Step #%d ====================================\n", i+1)
		fmt.Println(step.Description())
		err = step.Perform()
		if err != nil {
			break
		}
		taskprinter.New(ssp.task, logFile).Perform()
		fmt.Println(step.ResultMessage())
		if ssp.task.IsOptimalSolution {
			break
		}
	}
	return
}
