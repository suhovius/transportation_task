package main

func setPotential(targetList, sourceList *[]tableOuterCell, i, j int, cellCost float64) {
	(*targetList)[j].potential = cellCost - (*sourceList)[i].potential
	(*targetList)[j].isPotentialSet = true
}

func (t *Task) calculatePotentials() {
	// first potential is zero. U0= 0
	t.supplyList[0].isPotentialSet = true
	t.eachCell(
		func(i, j int) {
			cell := t.tableCells[i][j]
			if cell.deliveryAmount > 0 {
				switch {
				case t.supplyList[i].isPotentialSet:
					setPotential(&t.demandList, &t.supplyList, i, j, cell.cost)
				case t.demandList[j].isPotentialSet:
					setPotential(&t.supplyList, &t.demandList, j, i, cell.cost)
				}
			}
		},
	)
}
