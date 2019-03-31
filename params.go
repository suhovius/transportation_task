package main

import (
	"encoding/json"
)

// Params is a structure that contains API request parameters
type Params struct {
	SupplyList []int   `json:"supply_list"`
	DemandList []int   `json:"demand_list"`
	CostTable  [][]int `json:"cost_table"`
}

// ParseParams parses json and returns Params struct
func ParseParams(JSONData []byte) (parsedParams Params, err error) {
	err = json.Unmarshal(jsonBlob, &parsedParams)
	return
}
