package main

// IterationInitializer is a struct that implements AlgorithmStep interface
type IterationInitializer struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (ii *IterationInitializer) Description() string {
	return "Initialize task inner state before current iteration start"
}

// ResultMessage returns message about reults of step processing
func (ii *IterationInitializer) ResultMessage() string {
	return "Reset potentials, grades and circuit data"
}

// Perform cleans prevous changes at task's inner state
// to prepare it for new iteration
func (ii *IterationInitializer) Perform() (err error) {
	ii.resetPotentials()
	ii.resetGrades()
	ii.resetCircuit()
	return
}

func (ii *IterationInitializer) resetPotentials() {
	// reset demand potentials
	for i := range ii.task.DemandList {
		ii.task.DemandList[i].Potential = 0
		ii.task.DemandList[i].isPotentialSet = false
	}

	// reset supply potentials
	for i := range ii.task.SupplyList {
		ii.task.SupplyList[i].Potential = 0
		ii.task.SupplyList[i].isPotentialSet = false
	}
}

func (ii *IterationInitializer) resetGrades() {
	ii.task.minDeltaCell = cellIndexes{}
	ii.task.eachCell(
		func(i, j int) {
			ii.task.TableCells[i][j].delta = 0
		},
	)
}

func (ii *IterationInitializer) resetCircuit() {
	ii.task.thetaCell = PathVertex{}

	for _, vertex := range ii.task.path {
		ii.task.TableCells[vertex.i][vertex.j].sign = 0
	}

	ii.task.path = []PathVertex{}
}
