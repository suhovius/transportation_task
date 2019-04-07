package main

import (
	"errors"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
)

// DegeneracyChecker is a struct that implements AlgorithmStep interface
type DegeneracyChecker struct {
	AlgorithmStep
	task *taskmodel.Task
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

	dc.task.EachCell(func(i, j int) {
		if dc.task.TableCells[i][j].DeliveryAmount > 0 {
			cellsCount++
		}
	})
	return cellsCount
}

func (dc *DegeneracyChecker) basicCellsLimit() int {
	return len(dc.task.SupplyList) + len(dc.task.DemandList) - 1
}

func (dc *DegeneracyChecker) isDegenerate() bool {
	return dc.basicCellsCount() < dc.basicCellsLimit()
}
