package main

func (t *Task) calculateGrades() {
	t.eachCell(
		func(i, j int) {
			cP := &t.tableCells[i][j]
			if (*cP).deliveryAmount == 0 {
				(*cP).delta = (*cP).cost - t.supplyList[i].potential - t.demandList[j].potential
			}
		},
	)
}

func (t *Task) optimalSolutionCheck() (isValid bool) {
	t.calculateGrades()
	return
}
