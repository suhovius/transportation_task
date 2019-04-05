package main

import (
	"encoding/json"
)

// TaskParams is a structure that contains API request parameters
type TaskParams struct {
	SupplyList []int   `json:"supply_list"`
	DemandList []int   `json:"demand_list"`
	CostTable  [][]int `json:"cost_table"`
}

// ParseParams parses json and returns TaskParams struct
func ParseParams(jsonBlob []byte) (parsedParams TaskParams, err error) {
	err = json.Unmarshal(jsonBlob, &parsedParams)
	return
}
