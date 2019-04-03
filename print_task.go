package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// refactor it to taskPrinter{taskPtr: &task}.perform()

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
	return t.MinDeltaCell.isSet &&
		(t.MinDeltaCell.i == i && t.MinDeltaCell.j == j)
}

func (t *Task) minDeltaMarker(i, j int) string {
	if t.isMinDeltaCell(i, j) {
		return "\n(Min)"
	}
	return ""
}

func formatSign(sign rune) string {
	if sign > 0 {
		return fmt.Sprintf("\n(%s)", string(sign))
	}
	return ""
}

// Print prints current task processing state in the form of ASCII table
func (t *Task) Print() {
	data := make([][]string, len(t.supplyList))

	header := t.buildTableRow()
	header[0] = "→ Demand →\n----------\n↓ Supply ↓"
	for i, cell := range t.demandList {
		header[i+1] = fmt.Sprintf(
			"B[%d]= %d\nV[%d]= %d%s", i, roundToInt(cell.amount),
			i, roundToInt(cell.potential), fakeString(cell),
		)
	}

	for i, cellsRow := range t.tableCells {
		row := t.buildTableRow()
		supplier := t.supplyList[i]
		row[0] = fmt.Sprintf(
			"A[%d]= %d\nU[%d]= %d%s", i, roundToInt(supplier.amount),
			i, roundToInt(supplier.potential), fakeString(supplier),
		)
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf(
				"X= %d\nC= %d\nD= %d %s%s",
				roundToInt(cell.deliveryAmount), roundToInt(cell.cost),
				roundToInt(cell.delta), formatSign(cell.Sign),
				t.minDeltaMarker(i, j),
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
