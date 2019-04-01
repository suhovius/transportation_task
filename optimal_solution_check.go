package main

func (t *Task) calculateGrades() (hasNegativeValues bool) {
	var minDelta float64
	t.eachCell(
		func(i, j int) {
			cP := &t.tableCells[i][j]
			if (*cP).deliveryAmount == 0 {
				(*cP).delta =
					(*cP).cost - t.supplyList[i].potential - t.demandList[j].potential
				if (*cP).delta < 0 {
					hasNegativeValues = true
					if (*cP).delta < minDelta {
						t.minDeltaCell = cellIndexes{i: i, j: j}
					}
				}
			}
		},
	)
	return
}

func (t *Task) optimalSolutionCheck() (isOptimal bool) {
	return !t.calculateGrades()
}
