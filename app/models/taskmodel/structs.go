package taskmodel

import (
	"github.com/google/uuid"
)

// Task contains transportation task parameters and results
type Task struct {
	UUID              uuid.UUID
	SupplyList        []TableOuterCell `json:"supply_list"`
	DemandList        []TableOuterCell `json:"demand_list"`
	TableCells        [][]TableCell    `json:"table_cells"`
	MinDeltaCell      CellIndexes      `json:"-"`
	ThetaCell         PathVertex       `json:"-"` // TODO: Use CellIndexes later for this field
	Path              []PathVertex     `json:"-"`
	IsOptimalSolution bool             `json:"is_optimal_solution"`
	TotalDeliveryCost float64          `json:"-"`
}

// TableOuterCell defines table header and first column info
type TableOuterCell struct {
	Amount         float64 `json:"-"`
	Potential      float64 `json:"-"`
	IsPotentialSet bool    `json:"-"`
	IsFake         bool    `json:"is_fake"`
}

// TableCell defines table cell info
type TableCell struct {
	Cost           float64 `json:"-"`
	DeliveryAmount float64 `json:"-"`
	Delta          float64 `json:"-"`
	Sign           rune    `json:"-"` // + -
}

// CellIndexes is used to store indexes I, J of cells table matrix
type CellIndexes struct {
	I     int
	J     int
	IsSet bool
}

// PathVertex contains I, J indexes of the cycle path
type PathVertex struct {
	I, J int
}
