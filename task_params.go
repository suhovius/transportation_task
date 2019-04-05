package main

// TaskParams contains API request parameters
type TaskParams struct {
	SupplyList []int   `json:"supply_list"`
	DemandList []int   `json:"demand_list"`
	CostTable  [][]int `json:"cost_table"`
}
