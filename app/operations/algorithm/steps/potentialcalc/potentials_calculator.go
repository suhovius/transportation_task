package potentialcalc

import (
	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
)

// PotentialsCalculator is a struct that implements AlgorithmStep interface
type PotentialsCalculator struct {
	step.AlgorithmStep
	task *taskmodel.Task
}

// New returns new step instance
func New(task *taskmodel.Task) *PotentialsCalculator {
	return &PotentialsCalculator{task: task}
}

// Description returns step description info
func (pc *PotentialsCalculator) Description() string {
	return "Calculate Potentials"
}

// Task returns task pointer
func (pc *PotentialsCalculator) Task() *taskmodel.Task {
	return pc.task
}

// ResultMessage returns message about results of step processing
func (pc *PotentialsCalculator) ResultMessage() string {
	return "Potentials have been assigned to demand row and supply column"
}

// Perform implements step processing
func (pc *PotentialsCalculator) Perform() (err error) {
	// INFO: Potentials are nullified at IterationInitializer step
	// first Potential is zero. U0= 0
	t := pc.task
	t.SupplyList[0].IsPotentialSet = true
	t.EachCell(
		func(i, j int) {
			cell := t.TableCells[i][j]
			if cell.DeliveryAmount > 0 {
				switch {
				case t.SupplyList[i].IsPotentialSet:
					pc.setPotential(&t.DemandList, &t.SupplyList, i, j, cell.Cost)
				case t.DemandList[j].IsPotentialSet:
					pc.setPotential(&t.SupplyList, &t.DemandList, j, i, cell.Cost)
				}
			}
		},
	)

	return
}

func (pc *PotentialsCalculator) setPotential(
	targetList, sourceList *[]taskmodel.TableOuterCell, i, j int, cellCost float64,
) {
	(*targetList)[j].Potential = cellCost - (*sourceList)[i].Potential
	(*targetList)[j].IsPotentialSet = true
}
