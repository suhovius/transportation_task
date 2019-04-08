package taskform

// Params contains API request parameters
type Params struct {
	SupplyList []int   `json:"supply_list"`
	DemandList []int   `json:"demand_list"`
	CostTable  [][]int `json:"cost_table"`
}
