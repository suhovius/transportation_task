package main

// Task contains transportation task parameters and results
// TODO These fields should be Capitalized since they should be able to be
// converted to JSON so they should be public
type Task struct {
	supplyList   []tableOuterCell
	demandList   []tableOuterCell
	tableCells   [][]tableCell
	MinDeltaCell cellIndexes
	ThetaCell    PathVertex
	Path         []PathVertex
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
	isMinDelta     bool
	Sign           rune // + -
	// PathArrow      rune // ← ↑ → ↓ // TODO: Remove this field maybe
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
