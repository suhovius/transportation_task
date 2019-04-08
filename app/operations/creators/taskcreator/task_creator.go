package taskcreator

import (
	"bitbucket.org/suhovius/transportation_task/app/forms/taskform"
	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"github.com/google/uuid"
)

// TaskCreator provides logic for task creation
type TaskCreator struct {
	params *taskform.Params
}

// New returns new TaskCreator instance
func New(params *taskform.Params) *TaskCreator {
	return &TaskCreator{params: params}
}

// Perform creates task from params *TaskParams struct
func (tc *TaskCreator) Perform() (task taskmodel.Task) {
	task.UUID = uuid.New()
	task.SupplyList = make([]taskmodel.TableOuterCell, len(tc.params.SupplyList))
	for i, val := range tc.params.SupplyList {
		task.SupplyList[i] = taskmodel.TableOuterCell{Amount: float64(val)}
	}

	task.DemandList = make([]taskmodel.TableOuterCell, len(tc.params.DemandList))
	for i, val := range tc.params.DemandList {
		task.DemandList[i] = taskmodel.TableOuterCell{Amount: float64(val)}
	}

	task.TableCells = make([][]taskmodel.TableCell, len(tc.params.SupplyList))
	for i, row := range tc.params.CostTable {
		// assign table row
		task.TableCells[i] = make([]taskmodel.TableCell, len(row))
		for j, cost := range row {
			task.TableCells[i][j] = taskmodel.TableCell{Cost: float64(cost)}
		}
	}

	return
}
