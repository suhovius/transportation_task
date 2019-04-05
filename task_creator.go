package main

import (
	"github.com/google/uuid"
)

// TaskCreator provides logic for task creation
type TaskCreator struct {
	params *TaskParams
}

// Perform creates task from params *TaskParams struct
func (tc *TaskCreator) Perform() (task Task) {
	task.UUID = uuid.New()
	task.supplyList = make([]tableOuterCell, len(tc.params.SupplyList))
	for i, val := range tc.params.SupplyList {
		task.supplyList[i] = tableOuterCell{amount: float64(val)}
	}

	task.demandList = make([]tableOuterCell, len(tc.params.DemandList))
	for i, val := range tc.params.DemandList {
		task.demandList[i] = tableOuterCell{amount: float64(val)}
	}

	task.tableCells = make([][]tableCell, len(tc.params.SupplyList))
	for i, row := range tc.params.CostTable {
		// assign table row
		task.tableCells[i] = make([]tableCell, len(row))
		for j, cost := range row {
			task.tableCells[i][j] = tableCell{cost: float64(cost)}
		}
	}

	return
}
