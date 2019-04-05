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
		"Path: %v. Theta cell is at [%d][%d]",
		cb.task.path, cb.task.thetaCell.i, cb.task.thetaCell.j,
	)
}

func (cb *CircuitBuilder) lookForVertexWithMinDeliveryValue(pv *PathVertex) {
	if cb.thetaVertexPtr != nil {
		minAmount := cb.task.findCellByVertex(cb.thetaVertexPtr).DeliveryAmount
		newMinAmount := cb.task.findCellByVertex(pv).DeliveryAmount
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
		// find negative signed (-) cell with minimal delivery Amount
		cb.lookForVertexWithMinDeliveryValue(&vertex)
	}
	cb.task.TableCells[vertex.i][vertex.j].sign = sign
	return vertex
}

// Perform implements step processing
func (cb *CircuitBuilder) Perform() (err error) {
	err = cb.findPath()
	if err != nil {
		return err
	}
	return
}

func (cb *CircuitBuilder) findPath() (err error) {
	startVertex :=
		cb.addPathVertexWith(cb.task.minDeltaCell.i, cb.task.minDeltaCell.j)
	if !cb.searchHorizontally(startVertex) {
		// path has not been found
		return fmt.Errorf(
			"Can't find path for start vertex[%d][%d]",
			startVertex.i, startVertex.j,
		)
	}
	// path has been found
	cb.task.path = cb.path
	if cb.thetaVertexPtr != nil {
		cb.task.thetaCell = *cb.thetaVertexPtr
	} else {
		return fmt.Errorf(
			"Can't find path Theta cell for path %v", cb.task.path,
		)
	}

	return
}

func isNotCurrentCell(i1, i2 int) bool {
	return i1 != i2
}

// use pointer just to avoid copy of the whole TableCell structure
func isBasicCell(cell *TableCell) bool {
	// non zero delivery
	return cell.DeliveryAmount > 0
}

func (cb *CircuitBuilder) searchHorizontally(pv PathVertex) (isFound bool) {
	for j := 0; j < len(cb.task.DemandList); j++ {
		cellPtr := &cb.task.TableCells[pv.i][j]

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
	for i := 0; i < len(cb.task.SupplyList); i++ {
		cellPtr := &cb.task.TableCells[i][pv.j]

		if isNotCurrentCell(i, pv.i) && isBasicCell(cellPtr) {
			if cb.searchHorizontally(PathVertex{i: i, j: pv.j}) {
				cb.addPathVertexWith(i, pv.j)
				return true
			}
		}
	}
	return // false
}
