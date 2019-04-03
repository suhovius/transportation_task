package main

import "fmt"

// CycleBuilder is a struct that implements AlgorithmStep interface
type CycleBuilder struct {
	AlgorithmStep
	task *Task
	path []PathVertex
}

// Perform implements cycle
func (cb *CycleBuilder) Perform() (err error) {
	startVertex := PathVertex{
		i: cb.task.MinDeltaCell.i, j: cb.task.MinDeltaCell.j,
	}
	cb.path = append(cb.path, startVertex)
	if !cb.searchHorizontally(startVertex) {
		// path has not been found
		return fmt.Errorf(
			"Can't find path for start vertex[%d][%d]",
			startVertex.i, startVertex.j,
		)
	}
	// path has been found
	cb.task.Path = cb.path

	// cb.task
	// TODO: Starting from start point:
	// Find allowed connection points (according to the conditions)
	// Iterate over each allowed point to be connected
	// During each iteration call the same fuction recursively with this newly
	// selected point
	// until the cycle is built then stop
	// or stop when all variants have been checked including variants by range of row or column (
	// horizontal and vertical search
	//	to be defined how to check this:
	//  !!!it is easy to check when there is no any nearest vertexes
	//  which can satisfy the requirements
	// )
	return
}

func (cb *CycleBuilder) searchHorizontally(vertex PathVertex) (isFound bool) {
	return
}
