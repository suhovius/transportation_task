package main

import (
	"fmt"
)

// AmountDistributionChecker is a struct that implements AlgorithmStep interface
type AmountDistributionChecker struct {
	AlgorithmStep
	task *Task
}

// Description returns step description info
// TODO maybe this description should be moved to some different serivice object
func (adc *AmountDistributionChecker) Description() string {
	return "Perform amount distribution check"
}

// ResultMessage returns message about reults of step processing
func (adc *AmountDistributionChecker) ResultMessage() string {
	return "Success:\n - Sums of delivery amounts by columns and rows match each other"
}

// Perform implements supply amount redistribution step
func (adc *AmountDistributionChecker) Perform() (err error) {
	// check rows
	err = adc.amountDistributionCheckFor("supply", adc.rowCellFinder, adc.rowListsPair)

	if err != nil {
		return err
	}

	// check columns
	err = adc.amountDistributionCheckFor("demand", adc.colCellFinder, adc.colListsPair)

	// no any errors here
	return
}

func (adc *AmountDistributionChecker) buildError(
	entityName string, index int, actualSum, expectedSum float64,
) error {
	return fmt.Errorf(
		"Amount doesn't match: %s[%d] = %f Expected = %f",
		entityName, index, actualSum, expectedSum,
	)
}

func (adc *AmountDistributionChecker) sumMatchCheck(
	entityName string, index int, actualSum, expectedSum float64,
) (err error) {
	if !floatEquals(actualSum, expectedSum) {
		err = adc.buildError(entityName, index, actualSum, expectedSum)
	}
	return
}

func (adc *AmountDistributionChecker) amountDistributionCheckFor(
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

		err = adc.sumMatchCheck(entityName, i, actualSum, outerCell.amount)
		if err != nil {
			return err
		}
	}
	return
}

func (adc *AmountDistributionChecker) rowCellFinder(i, j int) tableCell {
	return adc.task.tableCells[i][j]
}

func (adc *AmountDistributionChecker) rowListsPair() (rowsList, colsList *[]tableOuterCell) {
	return &adc.task.supplyList, &adc.task.demandList
}

func (adc *AmountDistributionChecker) colCellFinder(i, j int) tableCell {
	// reverse indexes here
	return adc.task.tableCells[j][i]
}

func (adc *AmountDistributionChecker) colListsPair() (rowsList, colsList *[]tableOuterCell) {
	// reverse lists here
	return &adc.task.demandList, &adc.task.supplyList
}