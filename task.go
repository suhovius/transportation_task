package main

import (
	"github.com/google/uuid"
)

// Task contains transportation task parameters and results
// TODO These fields should be Capitalized since they should be able to be
// converted to JSON so they should be public
type Task struct {
	UUID              uuid.UUID
	SupplyList        []TableOuterCell `json:"supply_list"`
	DemandList        []TableOuterCell `json:"demand_list"`
	TableCells        [][]TableCell    `json:"table_cells"`
	minDeltaCell      cellIndexes
	thetaCell         PathVertex // TODO: Use cellIndexes later for this field
	path              []PathVertex
	IsOptimalSolution bool    `json:"is_optimal_solution"`
	TotalDeliveryCost float64 `json:"total_delivery_cost"`
}

// TableOuterCell defines table header and first column info
type TableOuterCell struct {
	Amount         float64 `json:"amount"`
	Potential      float64 `json:"potential"`
	isPotentialSet bool
	IsFake         bool `json:"is_fake"`
}

// TableCell defines table cell info
type TableCell struct {
	Cost           float64 `json:"cost"`
	DeliveryAmount float64 `json:"delivery_amount"`
	delta          float64 // These should be capitalized for JSON
	sign           rune    // + - // These should be capitalized for JSON
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

func (t *Task) findCellByVertex(pv *PathVertex) *TableCell {
	return &t.TableCells[pv.i][pv.j]
}

func (t *Task) eachCell(cellProcessor func(i, j int)) {
	for i, row := range t.TableCells {
		for j := range row {
			cellProcessor(i, j)
		}
	}
}
