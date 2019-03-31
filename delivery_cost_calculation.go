package main

func (c *tableCell) deliveryCost() float64 {
	return c.cost * c.deliveryAmount
}

func (t *Task) deliveryCost() float64 {
	var cost float64
	t.eachCell(func(i, j int) { cost += t.tableCells[i][j].deliveryCost() })
	return cost
}
