package taskprinter

import (
	"fmt"
	"os"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/utils/mathext"
	"github.com/olekukonko/tablewriter"
)

// TaskPrinter contains information for task printing
type TaskPrinter struct {
	task       *taskmodel.Task
	outputFile *os.File
}

// New returns new task printer
func New(task *taskmodel.Task, outputFile *os.File) *TaskPrinter {
	return &TaskPrinter{task: task, outputFile: outputFile}
}

func (p *TaskPrinter) buildTableRow() []string {
	return make([]string, len(p.task.DemandList)+1)
}

func fakeString(cell taskmodel.TableOuterCell) string {
	if cell.IsFake {
		return "\n(Fake)"
	}
	return ""
}

func (p *TaskPrinter) isMinDeltaCell(i, j int) bool {
	return p.task.MinDeltaCell.IsSet &&
		(p.task.MinDeltaCell.I == i && p.task.MinDeltaCell.J == j)
}

func (p *TaskPrinter) minDeltaMarker(i, j int) string {
	if p.isMinDeltaCell(i, j) {
		return "\nmin Δ"
	}
	return ""
}

func (p *TaskPrinter) thetaMarker(i, j int) string {
	// TODO: Fix zero value. Maybe Use CellIndexes type here
	if p.task.ThetaCell.I == i && p.task.ThetaCell.J == j {
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

// Perform prints current task processing state in the form of ASCII table
func (p *TaskPrinter) Perform() {
	t := p.task
	data := make([][]string, len(t.SupplyList))

	header := p.buildTableRow()
	header[0] = "→ Demand →\n----------\n↓ Supply ↓"
	for i, cell := range t.DemandList {
		header[i+1] = fmt.Sprintf(
			"B[%d]=%d\nV[%d]=%d%s", i, mathext.RoundToInt(cell.Amount),
			i, mathext.RoundToInt(cell.Potential), fakeString(cell),
		)
	}

	for i, cellsRow := range t.TableCells {
		row := p.buildTableRow()
		supplier := t.SupplyList[i]
		row[0] = fmt.Sprintf(
			"A[%d]=%d\nU[%d]=%d%s", i, mathext.RoundToInt(supplier.Amount),
			i, mathext.RoundToInt(supplier.Potential), fakeString(supplier),
		)
		for j, cell := range cellsRow {
			row[j+1] = fmt.Sprintf(
				"X=%d\n%s%s%s\n--------\nC=%d\nD=%d",
				mathext.RoundToInt(cell.DeliveryAmount),
				formatSign(cell.Sign),
				p.minDeltaMarker(i, j),
				p.thetaMarker(i, j),
				mathext.RoundToInt(cell.Cost),
				mathext.RoundToInt(cell.Delta),
			)
		}

		data[i] = row
	}

	table := tablewriter.NewWriter(p.outputFile)
	table.SetHeader(header)
	table.SetRowLine(true)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
