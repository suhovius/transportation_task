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
	task.SupplyList = make([]TableOuterCell, len(tc.params.SupplyList))
	for i, val := range tc.params.SupplyList {
		task.SupplyList[i] = TableOuterCell{Amount: float64(val)}
	}

	task.DemandList = make([]TableOuterCell, len(tc.params.DemandList))
	for i, val := range tc.params.DemandList {
		task.DemandList[i] = TableOuterCell{Amount: float64(val)}
	}

	task.TableCells = make([][]TableCell, len(tc.params.SupplyList))
	for i, row := range tc.params.CostTable {
		// assign table row
		task.TableCells[i] = make([]TableCell, len(row))
		for j, cost := range row {
			task.TableCells[i][j] = TableCell{Cost: float64(cost)}
		}
	}

	return
}
