package main

const elipsis = 0.001

func (t *Task) preventDegeneracy() {
	for i, cell := range t.demandList {
		t.demandList[i].amount = cell.amount + elipsis/float64(len(t.demandList))
	}
	t.supplyList[0].amount += elipsis
}

func (t *Task) basicCellsCount() int {
	cellsCount := 0

	t.eachCell(func(i, j int) {
		if t.tableCells[i][j].deliveryAmount > 0 {
			cellsCount++
		}
	})
	return cellsCount
}

func (t *Task) basicCellsLimit() int {
	return len(t.supplyList) + len(t.demandList) - 1
}

func (t *Task) isDegenerate() bool {
	return t.basicCellsCount() < t.basicCellsLimit()
}
