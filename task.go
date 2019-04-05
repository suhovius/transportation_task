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

func (t *Task) findCellByVertex(pv *PathVertex) *tableCell {
	return &t.tableCells[pv.i][pv.j]
}

func (t *Task) eachCell(cellProcessor func(i, j int)) {
	for i, row := range t.tableCells {
		for j := range row {
			cellProcessor(i, j)
		}
	}
}
