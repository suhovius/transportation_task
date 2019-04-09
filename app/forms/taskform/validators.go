package taskform

import (
	"errors"
	"fmt"
)

// Validate performs Params validation
func (p *Params) Validate() error {
	return p.performValidations(
		costTableRowsHasAtLeastOneRowValidator,
		supplyListSizeValidator,
		demandListSizeValidator,
		costTableRowsSizesValidator,
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

func supplyListSizeValidator(p *Params) (err error) {
	supplyCount := len(p.SupplyList)
	costRowsCount := len(p.CostTable)
	if supplyCount != costRowsCount {
		err = fmt.Errorf(
			"SupplyList size '%d' and CostTable rows count '%d' should be equal",
			supplyCount, costRowsCount,
		)
	}

	return
}

func demandListSizeValidator(p *Params) (err error) {
	demandCount := len(p.DemandList)
	columnsCount := len(p.CostTable[0])
	if demandCount != columnsCount {
		err = fmt.Errorf(
			"DemandList size '%d' and CostTable columns count '%d' should be equal",
			demandCount, columnsCount,
		)
	}

	return
}

func costTableRowsSizesValidator(p *Params) (err error) {
	demandCount := len(p.DemandList)
	for j, row := range p.CostTable {
		rowSize := len(row)
		if demandCount != rowSize {
			err = fmt.Errorf(
				"CostTable row [%d] size '%d' and DemandList size '%d' should be equal",
				j, rowSize, demandCount,
			)

			break
		}
	}
	return
}

func costTableRowsHasAtLeastOneRowValidator(p *Params) (err error) {
	if len(p.CostTable) == 0 {
		return errors.New("CostTable should have at least one row")
	}
	return
}
