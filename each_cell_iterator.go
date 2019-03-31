package main

func (t *Task) eachCell(cellProcessor func(i, j int)) {
	for i, row := range t.tableCells {
		for j := range row {
			cellProcessor(i, j)
		}
	}
}
