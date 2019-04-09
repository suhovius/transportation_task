package taskform

import (
	"fmt"
)

// Validate performs Params validation
func (p *Params) Validate() error {
	return p.performValidations(
		validateSupplyListSize,
		validateDemandListSize,
		validateCostTableRowsSizes,
	)
}

type validator func(p *Params) error

func (p *Params) performValidations(validators ...validator) (err error) {
	for _, v := range validators {
		if err = v(p); err != nil {
			break
		}
	}
	return
}

func validateSupplyListSize(p *Params) (err error) {
	supplyCount := len(p.SupplyList)
	costRowsCount := len(p.CostTable)
	if supplyCount != costRowsCount {
		err = fmt.Errorf(
			"Supply list size (%d) should be equal to Cost table rows count (%d)",
			supplyCount, costRowsCount,
		)
	}

	return
}

func validateDemandListSize(p *Params) (err error) {
	demandCount := len(p.DemandList)
	columnsCount := len(p.CostTable[0])
	if demandCount != columnsCount {
		err = fmt.Errorf(
			"Demand list size (%d) should be equal to Cost table columns count (%d)",
			demandCount, columnsCount,
		)
	}

	return
}

func validateCostTableRowsSizes(p *Params) (err error) {
	demandCount := len(p.DemandList)
	for j, row := range p.CostTable {
		rowSize := len(row)
		if demandCount != rowSize {
			err = fmt.Errorf(
				"Cost table row [%d] size (%d) should be equal to Demand list size (%d)",
				j, rowSize, demandCount,
			)

			break
		}
	}
	return
}
