package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// Task contains transportation task parameters and results
type Task struct {
	supplyList  []int
	demandList  []int
	costTable   [][]int
	resultTable [][]int
}

func (t Task) buildTableRow() []string {
	return make([]string, len(t.demandList)+1)
}

func (t Task) print() {
	data := make([][]string, len(t.supplyList))

	header := t.buildTableRow()
	header[0] = "Supply \\ Demand"
	for i, v := range t.demandList {
		header[i+1] = strconv.Itoa(v)
	}

	for i, costLine := range t.costTable {
		row := t.buildTableRow()
		row[0] = strconv.Itoa(t.supplyList[i])
		for j, v := range costLine {
			row[j+1] = fmt.Sprintf("X= %d / C= %d", t.resultTable[i][j], v)
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
