package sequence

import (
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
	"bitbucket.org/suhovius/transportation_task/app/operations/printers/taskprinter"
	log "github.com/sirupsen/logrus"
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

// RunWithLog starts all the AlgorithmStep handlers with logging of their results
func (ssp *StepsSequencePerformer) RunWithLog(le *log.Entry) (err error) {
	for i, step := range *ssp.steps {
		le.Infof("=== Step #%d ===", i+1)
		le.Info(step.Description())
		err = step.Perform()
		if err != nil {
			break
		}
		taskprinter.New(step.Task()).LogTaskState(le)
		le.Info(step.ResultMessage())
		if step.Task().IsOptimalSolution {
			break
		}
	}
	return
}
