package main

import "fmt"

// OptimalSolutionChecker is a struct that implements AlgorithmStep interface
type OptimalSolutionChecker struct {
	AlgorithmStep
	task      *Task
	isOptimal bool
}

// Description returns step description info
// TODO maybe this description should be moved to some different serivice object
func (osc *OptimalSolutionChecker) Description() string {
	return "Perform Optimal Solution Check"
}

// ResultMessage returns message about reults of step processing
func (osc *OptimalSolutionChecker) ResultMessage() string {
	var message string
	if osc.isOptimal {
		message = "Success: Solution is optimal.\nProccesing is Completed"
	} else {
		i := osc.task.MinDeltaCell.i
		j := osc.task.MinDeltaCell.j
		message = fmt.Sprintf(
			"Not Optimal:\n - Min Negative Delta Cell: D[%d][%d]= %d\n",
			i, j, roundToInt(osc.task.tableCells[i][j].delta),
		)
	}
	return message
}

// Perform implements step processing
func (osc *OptimalSolutionChecker) Perform() (err error) {
	osc.isOptimal = osc.verifyOptimality()

	return
}

func (osc *OptimalSolutionChecker) calculateGrades() (hasNegativeValues bool) {
	var minDelta float64
	t := osc.task
	t.eachCell(
		func(i, j int) {
			cP := &t.tableCells[i][j]
			if (*cP).deliveryAmount == 0 {
				(*cP).delta =
					(*cP).cost - t.supplyList[i].potential - t.demandList[j].potential
				if (*cP).delta < 0 {
					hasNegativeValues = true
					if (*cP).delta < minDelta {
						t.MinDeltaCell = cellIndexes{i: i, j: j, isSet: true}
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
