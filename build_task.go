package main

// Task contains transportation task parameters and results
type Task struct {
	supplyList []tableOuterCell
	demandList []tableOuterCell
	tableCells [][]tableCell
}

type tableOuterCell struct {
	amount         float64
	potential      float64
	isPotentialSet bool
	isFake         bool
}

type tableCell struct {
	cost           float64
	deliveryAmount float64
	delta          float64
	sign           rune
}

func buildTaskFromParams(params Params) Task {
	var task Task

	task.supplyList = make([]tableOuterCell, len(params.SupplyList))
	for i, val := range params.SupplyList {
		task.supplyList[i] = tableOuterCell{amount: float64(val)}
	}

	task.demandList = make([]tableOuterCell, len(params.DemandList))
	for i, val := range params.DemandList {
		task.demandList[i] = tableOuterCell{amount: float64(val)}
	}

	task.tableCells = make([][]tableCell, len(params.SupplyList))
	for i, row := range params.CostTable {
		// assign table row
		task.tableCells[i] = make([]tableCell, len(row))
		for j, cost := range row {
			task.tableCells[i][j] = tableCell{cost: float64(cost)}
		}
	}

	return task
}
