package main

import "fmt"

// StepsSequencePerformer contains array of AlgorithmStep handlers
// in the order of their start
type StepsSequencePerformer struct {
	steps *[]AlgorithmStep
	task  *Task
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
		// here also migth happen logging inside another service object wrapper
		ssp.task.Print()
		fmt.Println(step.ResultMessage())
		if ssp.task.IsOptimalSolution {
			break
		}
	}
	return
}
