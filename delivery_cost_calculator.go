package main

// DeliveryCostCalculator provides task delivery cost calculation logic
type DeliveryCostCalculator struct {
	task *Task
}

// Peform calculates task delivery cost
func (dcc *DeliveryCostCalculator) Peform() (cost float64) {
	dcc.task.eachCell(func(i, j int) {
		cell := dcc.task.tableCells[i][j]
		cost += cell.cost * cell.deliveryAmount
	})
	return cost
}
