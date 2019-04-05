package main

import "fmt"

const distVal = 0.001 // disturbance value

// DegeneracyPreventer is a struct that implements AlgorithmStep interface
type DegeneracyPreventer struct {
	AlgorithmStep
	task              *Task
	demandAmountsList []float64
}

// Description returns step description info
func (dp *DegeneracyPreventer) Description() string {
	return "Apply Degeneracy Prevention"
}

// ResultMessage returns message about reults of step processing
func (dp *DegeneracyPreventer) ResultMessage() string {
	return fmt.Sprintf(
		"Added %e to each demand Amount. Added %e to first supply Amount."+
			" Demand Amounts: %v"+
			" First supply Amount: %f",
		dp.demandDistrubance(), distVal, dp.demandAmountsList,
		dp.task.SupplyList[0].Amount,
	)
}

func (dp *DegeneracyPreventer) demandDistrubance() float64 {
	return distVal / float64(len(dp.task.DemandList))
}

// Perform implements step processing
func (dp *DegeneracyPreventer) Perform() (err error) {
	t := dp.task
	dp.demandAmountsList = make([]float64, len(t.DemandList))

	disturbance := dp.demandDistrubance()

	for i, cell := range t.DemandList {
		t.DemandList[i].Amount = cell.Amount + disturbance
		dp.demandAmountsList[i] = t.DemandList[i].Amount
	}
	t.SupplyList[0].Amount += distVal

	return
}
