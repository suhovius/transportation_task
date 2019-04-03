package main

import "fmt"

// CircuitBuilder is a struct that implements AlgorithmStep interface
type CircuitBuilder struct {
	AlgorithmStep
	task           *Task
	path           []PathVertex
	thetaVertexPtr *PathVertex
}

// Description returns step description info
func (cb *CircuitBuilder) Description() string {
	return "Finding the circuit path and minimal theta vertex"
}

// ResultMessage returns message about reults of step processing
func (cb *CircuitBuilder) ResultMessage() string {
	return fmt.Sprintf(
		"Success.\n - Path: %v\n - Theta cell is at [%d][%d]",
		cb.task.Path, cb.task.ThetaCell.i, cb.task.ThetaCell.j,
	)
}

func (cb *CircuitBuilder) lookForVertexWithMinDeliveryValue(pv *PathVertex) {
	if cb.thetaVertexPtr != nil {
		minAmount := cb.task.findCellByVertex(cb.thetaVertexPtr).deliveryAmount
		newMinAmount := cb.task.findCellByVertex(pv).deliveryAmount
		if minAmount > newMinAmount {
			// Smaller value have been found
			cb.thetaVertexPtr = pv
		}
	} else {
		// Set initial thetaVertexPtr value
		cb.thetaVertexPtr = pv
	}
}

func (cb *CircuitBuilder) addPathVertexWith(i, j int) PathVertex {
	vertex := PathVertex{i: i, j: j}
	cb.path = append(cb.path, vertex)
	var sign rune

	if len(cb.path)%2 != 0 {
		sign = '+'
	} else {
		sign = '-'
		// find negative signed (-) cell with minimal delivery amount
		cb.lookForVertexWithMinDeliveryValue(&vertex)
	}
	cb.task.tableCells[vertex.i][vertex.j].Sign = sign
	return vertex
}

// Perform implements Circuit building for transportation task solving
func (cb *CircuitBuilder) Perform() (err error) {
	err = cb.findPath()
	if err != nil {
		return err
	}
	return
}

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
	if cb.thetaVertexPtr != nil {
		cb.task.ThetaCell = *cb.thetaVertexPtr
	} else {
		return fmt.Errorf(
			"Can't find path Theta cell for path %v", cb.task.Path,
		)
	}

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
