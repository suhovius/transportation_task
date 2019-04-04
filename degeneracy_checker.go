package main

import "errors"

// DegeneracyChecker is a struct that implements AlgorithmStep interface
type DegeneracyChecker struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (dc *DegeneracyChecker) Description() string {
	return "Perform Degeneracy Check"
}

// ResultMessage returns message about reults of step processing
func (dc *DegeneracyChecker) ResultMessage() string {
	return "Solution is not Degenerate"
}

// Perform implements step processing
func (dc *DegeneracyChecker) Perform() (err error) {
	if dc.isDegenerate() {
		return errors.New("Degenerate Solution")
	}
	return
}

func (dc *DegeneracyChecker) basicCellsCount() int {
	cellsCount := 0

	dc.task.eachCell(func(i, j int) {
		if dc.task.tableCells[i][j].deliveryAmount > 0 {
			cellsCount++
		}
	})
	return cellsCount
}

func (dc *DegeneracyChecker) basicCellsLimit() int {
	return len(dc.task.supplyList) + len(dc.task.demandList) - 1
}

func (dc *DegeneracyChecker) isDegenerate() bool {
	return dc.basicCellsCount() < dc.basicCellsLimit()
}
