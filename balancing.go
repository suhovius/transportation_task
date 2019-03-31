package main

func listSum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func (t *Task) addFakeDemand(supplyExcess int) {
	// add fake demand value
	t.demandList = append(t.demandList, supplyExcess)
	for i := range t.supplyList {
		// set zero delivery cost for this fake demand
		t.costTable[i] = append(t.costTable[i], 0)
		// Keep proper dimensions for results table too
		t.resultTable[i] = append(t.resultTable[i], 0)
	}
}

func (t *Task) addFakeSupply(supplyDeficiency int) {
	// Add Fake Supplier
	t.supplyList = append(t.supplyList, supplyDeficiency)
	// Add row with zero price
	t.costTable = append(t.costTable, make([]int, len(t.demandList)))
	// Keep proper dimensions for results table too
	t.resultTable = append(t.resultTable, make([]int, len(t.demandList)))
}

func (t *Task) performBalancing() (kind string) {
	supplySumDiff := listSum(t.supplyList) - listSum(t.demandList)
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
