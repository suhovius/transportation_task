package taskmodel

import (
	"encoding/json"

	"bitbucket.org/suhovius/transportation_task/utils/mathext"
)

// TaskAlias is used to alter the way how Task is serialized to json
type TaskAlias Task

// AuxTask is used to alter the way how Task is serialized to json
type AuxTask struct {
	TotalDeliveryCost int `json:"total_delivery_cost"`
	*TaskAlias
}

// MarshalJSON provides custom way of json serialization for the Task struct
func (t *Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&AuxTask{
			TaskAlias:         (*TaskAlias)(t),
			TotalDeliveryCost: mathext.RoundToInt(t.TotalDeliveryCost),
		},
	)
}

// TableOuterCellAlias is used to alter the way how TableOuterCell is serialized to json
type TableOuterCellAlias TableOuterCell

// AuxTableOuterCell is used to alter the way how TableOuterCell is serialized to json
type AuxTableOuterCell struct {
	Amount    int `json:"amount"`
	Potential int `json:"potential"`
	*TableOuterCellAlias
}

// MarshalJSON provides custom way of json serialization for the TableOuterCell struct
func (oc *TableOuterCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&AuxTableOuterCell{
			TableOuterCellAlias: (*TableOuterCellAlias)(oc),
			Amount:              mathext.RoundToInt(oc.Amount),
			Potential:           mathext.RoundToInt(oc.Potential),
		},
	)
}

// TableCellAlias is used to alter the way how TableCell is serialized to json
type TableCellAlias TableCell

// AuxTableCell is used to alter the way how TableCell is serialized to json
type AuxTableCell struct {
	Cost           int `json:"cost"`
	DeliveryAmount int `json:"delivery_amount"`
	*TableCellAlias
}

// MarshalJSON provides custom way of json serialization for the TableCell struct
func (tc *TableCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&AuxTableCell{
			TableCellAlias: (*TableCellAlias)(tc),
			Cost:           mathext.RoundToInt(tc.Cost),
			DeliveryAmount: mathext.RoundToInt(tc.DeliveryAmount),
		},
	)
}
