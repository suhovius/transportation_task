package main

// NorthWestCornerSolutionFinder is a struct that implements AlgorithmStep interface
type NorthWestCornerSolutionFinder struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
func (nwcsf *NorthWestCornerSolutionFinder) Description() string {
	return "Calculate initial base plan with 'North West Corner' method"
}

// ResultMessage returns message about reults of step processing
func (nwcsf *NorthWestCornerSolutionFinder) ResultMessage() string {
	return "Done 'North West Corner' base plan calculation"
}

// Perform implements step processing
func (nwcsf *NorthWestCornerSolutionFinder) Perform() (err error) {
	task := nwcsf.task
	u := 0 // supplier index
	v := 0 // demand index
	// already supllied sums
	aS := make([]float64, len(task.DemandList))
	aD := make([]float64, len(task.SupplyList))

	for u < len(task.SupplyList) && v < len(task.DemandList) {
		if task.DemandList[v].Amount-aS[v] < task.SupplyList[u].Amount-aD[u] {
			// work with current row
			x := task.DemandList[v].Amount - aS[v]
			task.TableCells[u][v].DeliveryAmount = x
			aS[v] += x
			aD[u] += x
			v++
		} else {
			// go to the next row
			x := task.SupplyList[u].Amount - aD[u]
			task.TableCells[u][v].DeliveryAmount = x
			aS[v] += x
			aD[u] += x
			u++
		}
	}
	return
}
