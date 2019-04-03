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
	return fmt.Sprintf(
		"Perform Supply Redistribution with theta[%d][%d] value and signs (+) (-)",
		sr.task.ThetaCell.i, sr.task.ThetaCell.j,
	)
}

// Perform implements supply amount redistribution step
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
