package main

// Balancer is a struct that implements AlgorithmStep interface
type Balancer struct {
	AlgorithmStep
	task *Task
	kind string
}

// Description returns step description info
func (b *Balancer) Description() string {
	return "Perform Balancing"
}

// ResultMessage returns message about reults of step processing
func (b *Balancer) ResultMessage() (message string) {
	switch b.kind {
	case "nothing":
		message = "Balancing: Task is already balanced. Skip balancing"
	case "fake_demand":
		message = "Balancing: Add fake demand"
	case "fake_supply":
		message = "Balancing: Add fake supply"
	}
	return
}

// Perform implements step processing
func (b *Balancer) Perform() (err error) {
	supplySumDiff := b.listAmountSum(b.task.supplyList) - b.listAmountSum(b.task.demandList)
	switch {
	case supplySumDiff == 0:
		b.kind = "nothing"
	case supplySumDiff > 0:
		b.kind = "fake_demand"
		b.addFakeDemand(supplySumDiff)
	case supplySumDiff < 0:
		b.kind = "fake_supply"
		// convert negative value to positive
		b.addFakeSupply(-1 * supplySumDiff)
	}

	return
}

func (b *Balancer) listAmountSum(list []tableOuterCell) float64 {
	var sum float64
	for _, cell := range list {
		sum += cell.amount
	}
	return sum
}

func (b *Balancer) addFakeDemand(supplyExcess float64) {
	// add fake demand value
	b.task.demandList = append(
		b.task.demandList, tableOuterCell{amount: supplyExcess, isFake: true},
	)
	for i := range b.task.supplyList {
		// set zero delivery cost for this fake demand
		b.task.tableCells[i] = append(b.task.tableCells[i], tableCell{cost: 0})
	}
}

func (b *Balancer) addFakeSupply(supplyDeficiency float64) {
	// Add Fake Supplier
	b.task.supplyList = append(
		b.task.supplyList, tableOuterCell{amount: supplyDeficiency, isFake: true},
	)
	// Add row with zero price
	b.task.tableCells = append(
		b.task.tableCells, make([]tableCell, len(b.task.demandList)),
	)
}
