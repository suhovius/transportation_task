package main

func (t *Task) northWestCorner() {
	u := 0 // supplier index
	v := 0 // demand index
	// already supllied sums
	aS := make([]float64, len(t.DemandList))
	aD := make([]float64, len(t.SupplyList))

	for u < len(t.SupplyList) && v < len(t.DemandList) {
		if t.DemandList[v].Amount-aS[v] < t.SupplyList[u].Amount-aD[u] {
			// work with current row
			x := t.DemandList[v].Amount - aS[v]
			t.TableCells[u][v].DeliveryAmount = x
			aS[v] += x
			aD[u] += x
			v++
		} else {
			// go to the next row
			x := t.SupplyList[u].Amount - aD[u]
			t.TableCells[u][v].DeliveryAmount = x
			aS[v] += x
			aD[u] += x
			u++
		}
	}
}
