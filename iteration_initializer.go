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
	return "Reset potentials, grades."
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
	for i := range ii.task.demandList {
		ii.task.demandList[i].potential = 0
		ii.task.demandList[i].isPotentialSet = false
	}

	// reset supply potentials
	for i := range ii.task.supplyList {
		ii.task.supplyList[i].potential = 0
		ii.task.supplyList[i].isPotentialSet = false
	}
}

func (ii *IterationInitializer) resetGrades() {
	ii.task.MinDeltaCell = cellIndexes{}
	ii.task.eachCell(
		func(i, j int) {
			ii.task.tableCells[i][j].delta = 0
		},
	)
}

func (ii *IterationInitializer) resetCircuit() {
	ii.task.ThetaCell = PathVertex{}

	for _, vertex := range ii.task.Path {
		ii.task.tableCells[vertex.i][vertex.j].Sign = 0
	}

	ii.task.Path = []PathVertex{}
}