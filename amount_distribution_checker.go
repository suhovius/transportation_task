package main

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/utils/mathext"
)

// AmountDistributionChecker is a struct that implements AlgorithmStep interface
type AmountDistributionChecker struct {
	AlgorithmStep
	task *taskmodel.Task
}

// Description returns step description info
func (adc *AmountDistributionChecker) Description() string {
	return "Perform amount distribution check"
}

// ResultMessage returns message about reults of step processing
func (adc *AmountDistributionChecker) ResultMessage() string {
	return "Sums of delivery amounts by columns and rows match each other"
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
	if !mathext.FloatEquals(actualSum, expectedSum) {
		err = adc.buildError(entityName, index, actualSum, expectedSum)
	}
	return
}

func (adc *AmountDistributionChecker) amountDistributionCheckFor(
	entityName string,
	finder func(i, j int) taskmodel.TableCell,
	listsPair func() (r, c *[]taskmodel.TableOuterCell),
) (err error) {
	rowsList, colsList := listsPair()
	for i, outerCell := range *rowsList {
		var actualSum float64
		for j := range *colsList {
			actualSum += finder(i, j).DeliveryAmount
		}

		err = adc.sumMatchCheck(entityName, i, actualSum, outerCell.Amount)
		if err != nil {
			return err
		}
	}
	return
}

func (adc *AmountDistributionChecker) rowCellFinder(i, j int) taskmodel.TableCell {
	return adc.task.TableCells[i][j]
}

func (adc *AmountDistributionChecker) rowListsPair() (rowsList, colsList *[]taskmodel.TableOuterCell) {
	return &adc.task.SupplyList, &adc.task.DemandList
}

func (adc *AmountDistributionChecker) colCellFinder(i, j int) taskmodel.TableCell {
	// reverse indexes here
	return adc.task.TableCells[j][i]
}

func (adc *AmountDistributionChecker) colListsPair() (rowsList, colsList *[]taskmodel.TableOuterCell) {
	// reverse lists here
	return &adc.task.DemandList, &adc.task.SupplyList
}
