package taskmodel

// FindCellByVertex returns table cell by PathVertex I, J coordinates
func (t *Task) FindCellByVertex(pv *PathVertex) *TableCell {
	return &t.TableCells[pv.I][pv.J]
}

// EachCell iterates over task table and performs
// cellProcessor function for each cell
func (t *Task) EachCell(cellProcessor func(i, j int)) {
	for i, row := range t.TableCells {
		for j := range row {
			cellProcessor(i, j)
		}
	}
}
