package taskprinter

import (
	"fmt"
	"strings"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/utils/mathext"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

// TaskPrinter contains information for task printing
type TaskPrinter struct {
	task *taskmodel.Task
}

// New returns new task printer
func New(task *taskmodel.Task) *TaskPrinter {
	return &TaskPrinter{task: task}
}

// LogTaskState logs task state with logger
func (p *TaskPrinter) LogTaskState(le *log.Entry) {
	le.Infof("Current Task State Table:\n %s\n", p.RenderTableString())
}

// RenderTableString renders ASCII table to string
func (p *TaskPrinter) RenderTableString() string {
	td := p.prepareTableData()

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(td.header)
	table.SetRowLine(true)
	for _, v := range td.cells {
		table.Append(v)
	}
	table.Render()

	return tableString.String()
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

type tableData struct {
	header []string
	cells  [][]string
}

func (p *TaskPrinter) prepareTableData() *tableData {
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

	return &tableData{header: header, cells: data}
}
