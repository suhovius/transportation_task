package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func (t Task) buildTableRow() []string {
	return make([]string, len(t.demandList)+1)
}

func (t Task) print() {
	data := make([][]string, len(t.supplyList))

	header := t.buildTableRow()
	header[0] = "Supply \\ Demand"
	for i, cell := range t.demandList {
		header[i+1] = fmt.Sprintf("B= %d / V= %d", cell.amount, cell.potential)
	}

	for i, cellsRow := range t.tableCells {
		row := t.buildTableRow()
		supplier := t.supplyList[i]
		row[0] = fmt.Sprintf("B= %d / U= %d", supplier.amount, supplier.potential)
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf("X= %d / C= %d", cell.deliveryAmount, cell.cost)
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
