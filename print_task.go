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

func (t *Task) print() {
	data := make([][]string, len(t.supplyList))

	header := t.buildTableRow()
	header[0] = "Supply\n--------\nDemand"
	for i, cell := range t.demandList {
		header[i+1] = fmt.Sprintf("B= %f\nV= %f%s", cell.amount, cell.potential, fakeString(cell))
	}

	for i, cellsRow := range t.tableCells {
		row := t.buildTableRow()
		supplier := t.supplyList[i]
		row[0] = fmt.Sprintf("A= %f\nU= %f%s", supplier.amount, supplier.potential, fakeString(supplier))
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf("X= %f\nC= %f", cell.deliveryAmount, cell.cost)
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
