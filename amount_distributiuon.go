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

func (t *Task) amountDistributionCheck() (err error) {
	// check rows
	for i, supply := range t.supplyList {
		var actualSum float64
		for j := range t.demandList {
			actualSum += t.tableCells[i][j].deliveryAmount
		}

		err = sumMatchCheck("supply", i, actualSum, supply.amount)
		if err != nil {
			return err
		}
	}

	// check columns
	for i, demand := range t.demandList {
		var actualSum float64
		for j := range t.supplyList {
			actualSum += t.tableCells[j][i].deliveryAmount
		}

		err = sumMatchCheck("demand", i, actualSum, demand.amount)
		if err != nil {
			return err
		}
	}
	// no any errors here
	return
}
