package main

import "bitbucket.org/suhovius/transportation_task/app/models/taskmodel"

// DeliveryCostCalculator provides task delivery cost calculation logic
type DeliveryCostCalculator struct {
	task *taskmodel.Task
}

// Peform calculates task delivery cost
func (dcc *DeliveryCostCalculator) Peform() (cost float64) {
	dcc.task.EachCell(func(i, j int) {
		cell := dcc.task.TableCells[i][j]
		cost += cell.Cost * cell.DeliveryAmount
	})
	dcc.task.TotalDeliveryCost = cost
	return cost
}
