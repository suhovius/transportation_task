package main

import "bitbucket.org/suhovius/transportation_task/app/models/taskmodel"

// IterationInitializer is a struct that implements AlgorithmStep interface
type IterationInitializer struct {
	AlgorithmStep
	task *taskmodel.Task
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
		ii.task.DemandList[i].IsPotentialSet = false
	}

	// reset supply potentials
	for i := range ii.task.SupplyList {
		ii.task.SupplyList[i].Potential = 0
		ii.task.SupplyList[i].IsPotentialSet = false
	}
}

func (ii *IterationInitializer) resetGrades() {
	ii.task.MinDeltaCell = taskmodel.CellIndexes{}
	ii.task.EachCell(
		func(i, j int) {
			ii.task.TableCells[i][j].Delta = 0
		},
	)
}

func (ii *IterationInitializer) resetCircuit() {
	ii.task.ThetaCell = taskmodel.PathVertex{}

	for _, vertex := range ii.task.Path {
		ii.task.TableCells[vertex.I][vertex.J].Sign = 0
	}

	ii.task.Path = []taskmodel.PathVertex{}
}
