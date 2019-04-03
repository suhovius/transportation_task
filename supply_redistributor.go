package main

import "fmt"

// SupplyRedistributor is a struct that implements AlgorithmStep interface
type SupplyRedistributor struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
// TODO maybe this description should be moved to some different serivice object
func (sr *SupplyRedistributor) Description() string {
	return "Perform Supply Redistribution"
}

// ResultMessage returns message about reults of step processing
func (sr *SupplyRedistributor) ResultMessage() string {
	return fmt.Sprintf(
		"Delivery amounts have been updated according to theta[%d][%d] value and signs (+) (-)",
		sr.task.ThetaCell.i, sr.task.ThetaCell.j,
	)
}

// Perform implements step processing
func (sr *SupplyRedistributor) Perform() (err error) {
	thetaAmount := sr.task.findCellByVertex(&sr.task.ThetaCell).deliveryAmount
	for i, vertex := range sr.task.Path {
		cell := sr.task.findCellByVertex(&vertex)
		if i%2 == 0 {
			// even index has sign (+)
			cell.deliveryAmount += thetaAmount
		} else {
			// odd index has minus sign (-)
			cell.deliveryAmount -= thetaAmount
		}
	}
	if err != nil {
		return err
	}
	return
}
