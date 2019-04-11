package degeneracycheck

import (
	"errors"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
)

// DegeneracyChecker is a struct that implements AlgorithmStep interface
type DegeneracyChecker struct {
	step.AlgorithmStep
	task *taskmodel.Task
}

// New returns new step instance
func New(task *taskmodel.Task) *DegeneracyChecker {
	return &DegeneracyChecker{task: task}
}

// Task returns task pointer
func (dc *DegeneracyChecker) Task() *taskmodel.Task {
	return dc.task
}

// Description returns step description info
func (dc *DegeneracyChecker) Description() string {
	return "Perform Degeneracy Check"
}

// ResultMessage returns message about results of step processing
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
