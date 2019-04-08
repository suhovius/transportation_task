package sequence

import (
	"fmt"
	"os"

	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
	"bitbucket.org/suhovius/transportation_task/app/operations/printers/taskprinter"
)

// StepsSequencePerformer contains array of AlgorithmStep handlers
// in the order of their start
type StepsSequencePerformer struct {
	steps *[]step.AlgorithmStep
}

// New returns new sequence instance
func New(steps ...step.AlgorithmStep) *StepsSequencePerformer {
	return &StepsSequencePerformer{steps: &steps}
}

// Run starts all the AlgorithmStep handlers
// TODO: Refactor printing add logging at some separate object
func (ssp *StepsSequencePerformer) Run() (err error) {
	var logFile = os.Stdout // TODO: Use normal logger here or smth
	for i, step := range *ssp.steps {
		fmt.Printf("\n=== Step #%d ====================================\n", i+1)
		fmt.Println(step.Description())
		err = step.Perform()
		if err != nil {
			break
		}
		taskprinter.New(step.Task(), logFile).Perform()
		fmt.Println(step.ResultMessage())
		if step.Task().IsOptimalSolution {
			break
		}
	}
	return
}
