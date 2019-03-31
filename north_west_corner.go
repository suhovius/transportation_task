package main

func (t *Task) northWestCorner() {
	u := 0 // supplier index
	v := 0 // demand index
	// already supllied sums
	aS := make([]float64, len(t.demandList))
	aD := make([]float64, len(t.supplyList))

	for u < len(t.supplyList) && v < len(t.demandList) {
		if t.demandList[v].amount-aS[v] < t.supplyList[u].amount-aD[u] {
			// work with current row
			x := t.demandList[v].amount - aS[v]
			t.tableCells[u][v].deliveryAmount = x
			aS[v] += x
			aD[u] += x
			v++
		} else {
			// go to the next row
			x := t.supplyList[u].amount - aD[u]
			t.tableCells[u][v].deliveryAmount = x
			aS[v] += x
			aD[u] += x
			u++
		}
	}
}
