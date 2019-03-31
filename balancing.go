package main

func listAmountSum(list []tableOuterCell) int {
	sum := 0
	for _, cell := range list {
		sum += cell.amount
	}
	return sum
}

func (t *Task) addFakeDemand(supplyExcess int) {
	// add fake demand value
	t.demandList = append(t.demandList, tableOuterCell{amount: supplyExcess, isFake: true})
	for i := range t.supplyList {
		// set zero delivery cost for this fake demand
		t.tableCells[i] = append(t.tableCells[i], tableCell{cost: 0})
	}
}

func (t *Task) addFakeSupply(supplyDeficiency int) {
	// Add Fake Supplier
	t.supplyList = append(t.supplyList, tableOuterCell{amount: supplyDeficiency, isFake: true})
	// Add row with zero price
	t.tableCells = append(t.tableCells, make([]tableCell, len(t.demandList)))
}

func (t *Task) performBalancing() (kind string) {
	supplySumDiff := listAmountSum(t.supplyList) - listAmountSum(t.demandList)
	switch {
	case supplySumDiff == 0:
		kind = "nothing"
	case supplySumDiff > 0:
		kind = "fake_demand"
		t.addFakeDemand(supplySumDiff)
	case supplySumDiff < 0:
		kind = "fake_supply"
		// convert negative value to positive
		t.addFakeSupply(-1 * supplySumDiff)
	}
	return
}
