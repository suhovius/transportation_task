package main

func (c *tableCell) deliveryCost() int {
	return c.cost * c.deliveryAmount
}

func (t *Task) deliveryCost() int {
	cost := 0
	for _, row := range t.tableCells {
		for _, cell := range row {
			cost += cell.deliveryCost()
		}
	}
	return cost
}
