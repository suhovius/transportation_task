package main

import "fmt"

// SupplyRedistributor is a struct that implements AlgorithmStep interface
type SupplyRedistributor struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (sr *SupplyRedistributor) Description() string {
	return "Perform Supply Redistribution"
}

// ResultMessage returns message about reults of step processing
func (sr *SupplyRedistributor) ResultMessage() string {
	return fmt.Sprintf(
		"Delivery amounts have been updated according to theta[%d][%d] value and signs (+) (-)",
		sr.task.thetaCell.i, sr.task.thetaCell.j,
	)
}

// Perform implements step processing
func (sr *SupplyRedistributor) Perform() (err error) {
	thetaAmount := sr.task.findCellByVertex(&sr.task.thetaCell).DeliveryAmount
	for i, vertex := range sr.task.path {
		cell := sr.task.findCellByVertex(&vertex)
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
