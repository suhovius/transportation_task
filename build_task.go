package main

// Task contains transportation task parameters and results
type Task struct {
	supplyList []tableOuterCell
	demandList []tableOuterCell
	tableCells [][]tableCell
}

type tableCell struct {
	cost           int
	deliveryAmount int
	delta          int
	sign           rune
}

type tableOuterCell struct {
	value     int
	potential int
}

func buildTaskFromParams(params Params) Task {
	var task Task

	task.supplyList = make([]tableOuterCell, len(params.SupplyList))
	for i, val := range params.SupplyList {
		task.supplyList[i] = tableOuterCell{value: val}
	}

	task.demandList = make([]tableOuterCell, len(params.DemandList))
	for i, val := range params.DemandList {
		task.supplyList[i] = tableOuterCell{value: val}
	}

	task.tableCells = make([][]tableCell, len(params.SupplyList))
	for i, row := range params.CostTable {
		for j, value := range row {
			task.tableCells[i][j] = tableCell{cost: value}
		}
	}

	return task
}
