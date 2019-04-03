package main

import "fmt"

// CircuitBuilder is a struct that implements AlgorithmStep interface
type CircuitBuilder struct {
	AlgorithmStep
	task *Task
	path []PathVertex
}

func (cb *CircuitBuilder) addPathVertexWith(i, j int) PathVertex {
	vertex := PathVertex{i: i, j: j}
	cb.path = append(cb.path, vertex)
	cb.task.tableCells[vertex.i][vertex.j].PathArrow = '*' // TODO: This might be boolean
	return vertex
}

// Perform implements Circuit building for transportation task solving
func (cb *CircuitBuilder) Perform() (err error) {
	err = cb.findPath()
	if err != nil {
		return err
	}
	// cb.markPathAtCells()
	return
}

// func (cb *CircuitBuilder) markPathAtCells() {
// 	var arrow rune
// 	for _, vertex := range cb.path {
// 		arrow = '*'
// 		// TODO: Set arrows here by checking the prev and current index which is bigger
// 		// this will let to define the direction
// 		// if i == 0 {
// 		// 	arrow = '*'
// 		// } else {
// 		// 	if i < len(cb.path) - 1 {
// 		// 		// horizontal line
// 		// 		if cp.path.i == cp.path[i+1].i {

// 		// 		}
// 		// 		arrow = '←' // ← ↑ → ↓
// 		// 	}

// 		// }
// 		cb.task.tableCells[vertex.i][vertex.j].PathArrow = arrow
// 	}
// }

func (cb *CircuitBuilder) findPath() (err error) {
	startVertex :=
		cb.addPathVertexWith(cb.task.MinDeltaCell.i, cb.task.MinDeltaCell.j)
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
	// until the Circuit is built then stop
	// or stop when all variants have been checked including variants by range of row or column (
	// horizontal and vertical search
	//	to be defined how to check this:
	//  !!!it is easy to check when there is no any nearest vertexes
	//  which can satisfy the requirements
	// )
	return
}

func isNotCurrentCell(i1, i2 int) bool {
	return i1 != i2
}

// use pointer just to avoid copy of the whole tableCell structure
func isBasicCell(cell *tableCell) bool {
	// non zero delivery
	return cell.deliveryAmount > 0
}

func (cb *CircuitBuilder) searchHorizontally(pv PathVertex) (isFound bool) {
	for j := 0; j < len(cb.task.demandList); j++ {
		cellPtr := &cb.task.tableCells[pv.i][j]

		if isNotCurrentCell(j, pv.j) && isBasicCell(cellPtr) {
			// if we can connect with start vertex, then path is completed
			if j == cb.path[0].j {
				cb.addPathVertexWith(pv.i, j)
				return true // Circuit completed
			}

			if cb.searchVertically(PathVertex{i: pv.i, j: j}) {
				cb.addPathVertexWith(pv.i, j)
				return true // Circuit completed
			}
		}
	}
	return // false
}

func (cb *CircuitBuilder) searchVertically(pv PathVertex) (isFound bool) {
	for i := 0; i < len(cb.task.supplyList); i++ {
		cellPtr := &cb.task.tableCells[i][pv.j]

		if isNotCurrentCell(i, pv.i) && isBasicCell(cellPtr) {
			if cb.searchHorizontally(PathVertex{i: i, j: pv.j}) {
				cb.addPathVertexWith(i, pv.j)
				return true
			}
		}
	}
	return // false
}
