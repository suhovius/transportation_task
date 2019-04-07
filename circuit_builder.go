package main

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
)

// CircuitBuilder is a struct that implements AlgorithmStep interface
type CircuitBuilder struct {
	AlgorithmStep
	task           *taskmodel.Task
	path           []taskmodel.PathVertex
	thetaVertexPtr *taskmodel.PathVertex
}

// Description returns step description info
func (cb *CircuitBuilder) Description() string {
	return "Finding the circuit path and minimal theta vertex"
}

// ResultMessage returns message about reults of step processing
func (cb *CircuitBuilder) ResultMessage() string {
	return fmt.Sprintf(
		"Path: %v. Theta cell is at [%d][%d]",
		cb.task.Path, cb.task.ThetaCell.I, cb.task.ThetaCell.J,
	)
}

func (cb *CircuitBuilder) lookForVertexWithMinDeliveryValue(pv *taskmodel.PathVertex) {
	if cb.thetaVertexPtr != nil {
		minAmount := cb.task.FindCellByVertex(cb.thetaVertexPtr).DeliveryAmount
		newMinAmount := cb.task.FindCellByVertex(pv).DeliveryAmount
		if minAmount > newMinAmount {
			// Smaller value have been found
			cb.thetaVertexPtr = pv
		}
	} else {
		// Set initial thetaVertexPtr value
		cb.thetaVertexPtr = pv
	}
}

func (cb *CircuitBuilder) addPathVertexWith(i, j int) taskmodel.PathVertex {
	vertex := taskmodel.PathVertex{I: i, J: j}
	cb.path = append(cb.path, vertex)
	var sign rune

	if len(cb.path)%2 != 0 {
		sign = '+'
	} else {
		sign = '-'
		// find negative signed (-) cell with minimal delivery Amount
		cb.lookForVertexWithMinDeliveryValue(&vertex)
	}
	cb.task.TableCells[vertex.I][vertex.J].Sign = sign
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
		cb.addPathVertexWith(cb.task.MinDeltaCell.I, cb.task.MinDeltaCell.J)
	if !cb.searchHorizontally(startVertex) {
		// path has not been found
		return fmt.Errorf(
			"Can't find path for start vertex[%d][%d]",
			startVertex.I, startVertex.J,
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

// use pointer just to avoid copy of the whole TableCell structure
func isBasicCell(cell *taskmodel.TableCell) bool {
	// non zero delivery
	return cell.DeliveryAmount > 0
}

func (cb *CircuitBuilder) searchHorizontally(pv taskmodel.PathVertex) (isFound bool) {
	for j := 0; j < len(cb.task.DemandList); j++ {
		cellPtr := &cb.task.TableCells[pv.I][j]

		if isNotCurrentCell(j, pv.J) && isBasicCell(cellPtr) {
			// if we can connect with start vertex, then path is completed
			if j == cb.path[0].J {
				cb.addPathVertexWith(pv.I, j)
				return true // Circuit completed
			}

			if cb.searchVertically(taskmodel.PathVertex{I: pv.I, J: j}) {
				cb.addPathVertexWith(pv.I, j)
				return true // Circuit completed
			}
		}
	}
	return // false
}

func (cb *CircuitBuilder) searchVertically(pv taskmodel.PathVertex) (isFound bool) {
	for i := 0; i < len(cb.task.SupplyList); i++ {
		cellPtr := &cb.task.TableCells[i][pv.J]

		if isNotCurrentCell(i, pv.I) && isBasicCell(cellPtr) {
			if cb.searchHorizontally(taskmodel.PathVertex{I: i, J: pv.J}) {
				cb.addPathVertexWith(i, pv.J)
				return true
			}
		}
	}
	return // false
}
