package supplyredistrib

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
)

// SupplyRedistributor is a struct that implements AlgorithmStep interface
type SupplyRedistributor struct {
	step.AlgorithmStep
	task *taskmodel.Task
}

// New returns new step instance
func New(task *taskmodel.Task) *SupplyRedistributor {
	return &SupplyRedistributor{task: task}
}

// Task returns task pointer
func (sr *SupplyRedistributor) Task() *taskmodel.Task {
	return sr.task
}

// Description returns step description info
func (sr *SupplyRedistributor) Description() string {
	return "Perform Supply Redistribution"
}

// ResultMessage returns message about reults of step processing
func (sr *SupplyRedistributor) ResultMessage() string {
	return fmt.Sprintf(
		"Delivery amounts have been updated according to theta[%d][%d] value and signs (+) (-)",
		sr.task.ThetaCell.I, sr.task.ThetaCell.J,
	)
}

// Perform implements step processing
func (sr *SupplyRedistributor) Perform() (err error) {
	thetaAmount := sr.task.FindCellByVertex(&sr.task.ThetaCell).DeliveryAmount
	for i, vertex := range sr.task.Path {
		cell := sr.task.FindCellByVertex(&vertex)
		if i%2 == 0 {
			// even index has sign (+)
			cell.DeliveryAmount += thetaAmount
		} else {
			// odd index has minus sign (-)
			cell.DeliveryAmount -= thetaAmount
		}
	}
	if err != nil {
		return err
	}
	return
}
