package taskmodel

// FindCellByVertex returns table cell by PathVertex I, J coordinates
func (t *Task) FindCellByVertex(pv *PathVertex) *TableCell {
	return t.FindCellByIndexes(pv.I, pv.J)
}

// FindCellByIndexes returns table cell by i, j indexes
func (t *Task) FindCellByIndexes(i, j int) *TableCell {
	return &t.TableCells[i][j]
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
