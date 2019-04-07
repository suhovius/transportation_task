package main

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/utils/mathext"
)

// OptimalSolutionChecker is a struct that implements AlgorithmStep interface
type OptimalSolutionChecker struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (osc *OptimalSolutionChecker) Description() string {
	return "Perform Optimal Solution Check"
}

// ResultMessage returns message about reults of step processing
func (osc *OptimalSolutionChecker) ResultMessage() string {
	var message string
	if osc.task.IsOptimalSolution {
		message = "Solution is optimal. Proccesing is Completed"
	} else {
		i := osc.task.minDeltaCell.i
		j := osc.task.minDeltaCell.j
		message = fmt.Sprintf(
			"Not Optimal Solution. Min Negative Delta Cell: D[%d][%d]= %d\n",
			i, j, mathext.RoundToInt(osc.task.TableCells[i][j].delta),
		)
	}
	return message
}

// Perform implements step processing
func (osc *OptimalSolutionChecker) Perform() (err error) {
	osc.task.IsOptimalSolution = osc.verifyOptimality()

	return
}

func (osc *OptimalSolutionChecker) calculateGrades() (hasNegativeValues bool) {
	var minDelta float64
	t := osc.task
	t.eachCell(
		func(i, j int) {
			cP := &t.TableCells[i][j]
			if (*cP).DeliveryAmount == 0 {
				(*cP).delta =
					(*cP).Cost - t.SupplyList[i].Potential - t.DemandList[j].Potential
				if (*cP).delta < 0 {
					hasNegativeValues = true
					if (*cP).delta < minDelta {
						t.minDeltaCell = cellIndexes{i: i, j: j, isSet: true}
					}
				}
			}
		},
	)
	return
}

func (osc *OptimalSolutionChecker) verifyOptimality() bool {
	return !osc.calculateGrades()
}
