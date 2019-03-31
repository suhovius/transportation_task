package main

func (c *tableCell) deliveryCost() float64 {
	return c.cost * c.deliveryAmount
}

func (t *Task) deliveryCost() float64 {
	var cost float64
	for _, row := range t.tableCells {
		for _, cell := range row {
			cost += cell.deliveryCost()
		}
	}
	return cost
}
