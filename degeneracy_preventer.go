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
		"Added %e to each demand amount. Added %e to first supply amount."+
			" Demand amounts: %v"+
			" First supply amount: %f",
		dp.demandDistrubance(), distVal, dp.demandAmountsList,
		dp.task.supplyList[0].amount,
	)
}

func (dp *DegeneracyPreventer) demandDistrubance() float64 {
	return distVal / float64(len(dp.task.demandList))
}

// Perform implements step processing
func (dp *DegeneracyPreventer) Perform() (err error) {
	t := dp.task
	dp.demandAmountsList = make([]float64, len(t.demandList))

	disturbance := dp.demandDistrubance()

	for i, cell := range t.demandList {
		t.demandList[i].amount = cell.amount + disturbance
		dp.demandAmountsList[i] = t.demandList[i].amount
	}
	t.supplyList[0].amount += distVal

	return
}
