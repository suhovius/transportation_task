package main

// SupplyRedistributor is a struct that implements AlgorithmStep interface
type SupplyRedistributor struct {
	AlgorithmStep
	task *Task
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
