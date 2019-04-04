package main

// Task contains transportation task parameters and results
// TODO These fields should be Capitalized since they should be able to be
// converted to JSON so they should be public
type Task struct {
	supplyList        []tableOuterCell // These should be capitalized for JSON
	demandList        []tableOuterCell // These should be capitalized for JSON
	tableCells        [][]tableCell    // These should be capitalized for JSON
	MinDeltaCell      cellIndexes
	ThetaCell         PathVertex // Maybe Use cellIndexes type here
	Path              []PathVertex
	IsOptimalSolution bool
}

type tableOuterCell struct {
	amount         float64 // These should be capitalized for JSON
	potential      float64 // These should be capitalized for JSON
	isPotentialSet bool
	isFake         bool // These should be capitalized for JSON
}

type tableCell struct {
	cost           float64 // These should be capitalized for JSON
	deliveryAmount float64 // These should be capitalized for JSON
	delta          float64 // These should be capitalized for JSON
	isMinDelta     bool
	Sign           rune // + - // These should be capitalized for JSON
}

type cellIndexes struct {
	i     int
	j     int
	isSet bool
}

// PathVertex contains i, j indexes of the cycle path
type PathVertex struct {
	i, j int
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

func (task *Task) findCellByVertex(pv *PathVertex) *tableCell {
	return &task.tableCells[pv.i][pv.j]
}
