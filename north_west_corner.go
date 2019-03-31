package main

func (t *Task) northWestCorner() {
	u := 0 // supplier index
	v := 0 // demand index
	// already supllied sums
	aS := make([]int, len(t.demandList))
	aD := make([]int, len(t.supplyList))

	for u < len(t.supplyList) && v < len(t.demandList) {
		if t.demandList[v]-aS[v] < t.supplyList[u]-aD[u] {
			// work with current row
			x := t.demandList[v] - aS[v]
			t.resultTable[u][v] = x
			aS[v] += x
			aD[u] += x
			v++
		} else {
			// go to the next row
			x := t.supplyList[u] - aD[u]
			t.resultTable[u][v] = x
			aS[v] += x
			aD[u] += x
			u++
		}
	}
}
