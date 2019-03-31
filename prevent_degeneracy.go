package main

const elipsis = 0.001

func (t *Task) preventDegeneracy() {
	for i, cell := range t.demandList {
		t.demandList[i].amount = cell.amount + elipsis/float64(len(t.demandList))
	}
	t.supplyList[0].amount += elipsis
}
