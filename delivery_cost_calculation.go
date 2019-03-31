package main

func (t Task) deliveryCost() int {
	cost := 0
	for i, costRow := range t.costTable {
		for j, costValue := range costRow {
			cost += costValue * t.resultTable[i][j]
		}
	}
	return cost
}
