package main

import (
	"fmt"
)

func buildError(entityName string, index int, actualSum, expectedSum float64) error {
	return fmt.Errorf(
		"Amount doesn't match: %s[%d] = %f Expected = %f", entityName, index, actualSum, expectedSum,
	)
}

func sumMatchCheck(entityName string, index int, actualSum, expectedSum float64) (err error) {
	if !floatEquals(actualSum, expectedSum) {
		err = buildError(entityName, index, actualSum, expectedSum)
	}
	return
}

func (t *Task) amountDistributionCheckFor(
	entityName string,
	finder func(i, j int) tableCell,
	listsPair func() (r, c *[]tableOuterCell),
) (err error) {
	rowsList, colsList := listsPair()
	for i, outerCell := range *rowsList {
		var actualSum float64
		for j := range *colsList {
			actualSum += finder(i, j).deliveryAmount
		}

		err = sumMatchCheck(entityName, i, actualSum, outerCell.amount)
		if err != nil {
			return err
		}
	}
	return
}

func (t *Task) rowCellFinder(i, j int) tableCell {
	return t.tableCells[i][j]
}

func (t *Task) rowListsPair() (rowsList, colsList *[]tableOuterCell) {
	return &t.supplyList, &t.demandList
}

func (t *Task) colCellFinder(i, j int) tableCell {
	// reverse indexes here
	return t.tableCells[j][i]
}

func (t *Task) colListsPair() (rowsList, colsList *[]tableOuterCell) {
	// reverse lists here
	return &t.demandList, &t.supplyList
}

func (t *Task) amountDistributionCheck() (err error) {
	// check rows
	err = t.amountDistributionCheckFor("supply", t.rowCellFinder, t.rowListsPair)

	if err != nil {
		return err
	}

	// check columns
	err = t.amountDistributionCheckFor("demand", t.colCellFinder, t.colListsPair)

	// no any errors here
	return
}
