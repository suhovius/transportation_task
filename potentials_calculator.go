package main

// PotentialsCalculator is a struct that implements AlgorithmStep interface
type PotentialsCalculator struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (pc *PotentialsCalculator) Description() string {
	return "Calculate Potentials"
}

// ResultMessage returns message about reults of step processing
func (pc *PotentialsCalculator) ResultMessage() string {
	return "Potentials have been assigned to demand row and supply column"
}

// Perform implements step processing
func (pc *PotentialsCalculator) Perform() (err error) {
	// Info potentials are nullified at IterationInitializer step
	// first potential is zero. U0= 0
	t := pc.task
	t.supplyList[0].isPotentialSet = true
	t.eachCell(
		func(i, j int) {
			cell := t.tableCells[i][j]
			if cell.deliveryAmount > 0 {
				switch {
				case t.supplyList[i].isPotentialSet:
					pc.setPotential(&t.demandList, &t.supplyList, i, j, cell.cost)
				case t.demandList[j].isPotentialSet:
					pc.setPotential(&t.supplyList, &t.demandList, j, i, cell.cost)
				}
			}
		},
	)

	return
}

func (pc *PotentialsCalculator) setPotential(
	targetList, sourceList *[]tableOuterCell, i, j int, cellCost float64,
) {
	(*targetList)[j].potential = cellCost - (*sourceList)[i].potential
	(*targetList)[j].isPotentialSet = true
}
