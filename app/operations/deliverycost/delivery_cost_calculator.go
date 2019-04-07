package deliverycost

import "bitbucket.org/suhovius/transportation_task/app/models/taskmodel"

// Calculator provides task delivery cost calculation logic
type Calculator struct {
	task *taskmodel.Task
}

// New returns new instance
func New(task *taskmodel.Task) *Calculator {
	return &Calculator{task: task}
}

// Perform calculates task delivery cost
func (dcc *Calculator) Perform() (cost float64) {
	dcc.task.EachCell(func(i, j int) {
		cell := dcc.task.TableCells[i][j]
		cost += cell.Cost * cell.DeliveryAmount
	})
	dcc.task.TotalDeliveryCost = cost
	return cost
}
