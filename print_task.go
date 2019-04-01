package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func (t *Task) buildTableRow() []string {
	return make([]string, len(t.demandList)+1)
}

func fakeString(cell tableOuterCell) string {
	if cell.isFake {
		return "\n(Fake)"
	}
	return ""
}

func (t *Task) isMinDeltaCell(i, j int) bool {
	return t.minDeltaCell.isSet &&
		(t.minDeltaCell.i == i && t.minDeltaCell.j == j)
}

func (t *Task) minDeltaMarker(i, j int) string {
	if t.isMinDeltaCell(i, j) {
		return "\n(Min)"
	}
	return ""
}

// Print prints current task processing state in the form of ASCII table
func (t *Task) Print() {
	data := make([][]string, len(t.supplyList))

	header := t.buildTableRow()
	header[0] = "→ Supply →\n----------\n↓ Demand ↓"
	for i, cell := range t.demandList {
		header[i+1] = fmt.Sprintf(
			"B[%d]= %f\nV[%d]= %f%s", i, cell.amount,
			i, cell.potential, fakeString(cell),
		)
	}

	for i, cellsRow := range t.tableCells {
		row := t.buildTableRow()
		supplier := t.supplyList[i]
		row[0] = fmt.Sprintf(
			"A[%d]= %f\nU[%d]= %f%s", i, supplier.amount,
			i, supplier.potential, fakeString(supplier),
		)
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf(
				"X= %f\nC= %f\nD= %f%s", cell.deliveryAmount, cell.cost,
				cell.delta, t.minDeltaMarker(i, j),
			)
		}

		data[i] = row
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
