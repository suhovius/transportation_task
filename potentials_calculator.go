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
	// Info Potentials are nullified at IterationInitializer step
	// first Potential is zero. U0= 0
	t := pc.task
	t.SupplyList[0].isPotentialSet = true
	t.eachCell(
		func(i, j int) {
			cell := t.TableCells[i][j]
			if cell.DeliveryAmount > 0 {
				switch {
				case t.SupplyList[i].isPotentialSet:
					pc.setPotential(&t.DemandList, &t.SupplyList, i, j, cell.Cost)
				case t.DemandList[j].isPotentialSet:
					pc.setPotential(&t.SupplyList, &t.DemandList, j, i, cell.Cost)
				}
			}
		},
	)

	return
}

func (pc *PotentialsCalculator) setPotential(
	targetList, sourceList *[]TableOuterCell, i, j int, cellCost float64,
) {
	(*targetList)[j].Potential = cellCost - (*sourceList)[i].Potential
	(*targetList)[j].isPotentialSet = true
}
