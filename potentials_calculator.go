package main

// PotentialsCalculator is a struct that implements AlgorithmStep interface
type PotentialsCalculator struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
// TODO maybe this description should be moved to some different serivice object
func (pc *PotentialsCalculator) Description() string {
	return "Calculate Potentials"
}

// ResultMessage returns message about reults of step processing
func (pc *PotentialsCalculator) ResultMessage() string {
	return "Success:\n - Potentials have been assigned to demand row and supply column"
}

// Perform implements step processing
func (pc *PotentialsCalculator) Perform() (err error) {
	// TODO nullify previously set potentials from previous step
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
