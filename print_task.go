package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// TODO: Refactor it to TaskPrinter{task: &task}.Perform()

func (t *Task) buildTableRow() []string {
	return make([]string, len(t.DemandList)+1)
}

func fakeString(cell TableOuterCell) string {
	if cell.IsFake {
		return "\n(Fake)"
	}
	return ""
}

func (t *Task) isminDeltaCell(i, j int) bool {
	return t.minDeltaCell.isSet &&
		(t.minDeltaCell.i == i && t.minDeltaCell.j == j)
}

func (t *Task) mindeltaMarker(i, j int) string {
	if t.isminDeltaCell(i, j) {
		return "\nmin Δ"
	}
	return ""
}

func thetaMarker(t *Task, i, j int) string {
	// TODO: Fix zero value. Maybe Use cellIndexes type here
	if t.thetaCell.i == i && t.thetaCell.j == j {
		return "\nmin θ"
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
	data := make([][]string, len(t.SupplyList))

	header := t.buildTableRow()
	header[0] = "→ Demand →\n----------\n↓ Supply ↓"
	for i, cell := range t.DemandList {
		header[i+1] = fmt.Sprintf(
			"B[%d]=%d\nV[%d]=%d%s", i, roundToInt(cell.Amount),
			i, roundToInt(cell.Potential), fakeString(cell),
		)
	}

	for i, cellsRow := range t.TableCells {
		row := t.buildTableRow()
		supplier := t.SupplyList[i]
		row[0] = fmt.Sprintf(
			"A[%d]=%d\nU[%d]=%d%s", i, roundToInt(supplier.Amount),
			i, roundToInt(supplier.Potential), fakeString(supplier),
		)
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf(
				"X=%d\n%s%s%s\n--------\nC=%d\nD=%d",
				roundToInt(cell.DeliveryAmount),
				formatSign(cell.sign),
				t.mindeltaMarker(i, j),
				thetaMarker(t, i, j),
				roundToInt(cell.Cost),
				roundToInt(cell.delta),
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
