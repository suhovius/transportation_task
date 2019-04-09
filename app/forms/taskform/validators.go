package taskform

import (
	"fmt"

	"bitbucket.org/suhovius/transportation_task/app/forms/validate"
)

type supplyListValidator struct {
	validate.Validator
	params *Params
}

func (v *supplyListValidator) Perform() (err error) {
	supplyCount := len(v.params.SupplyList)
	costRowsCount := len(v.params.CostTable)
	if supplyCount != costRowsCount {
		err = fmt.Errorf(
			"Supply list size %d should be equal to Cost table rows count %d",
			supplyCount, costRowsCount,
		)
	}

	return
}

// Validate performs Params validation
func (p *Params) Validate() error {
	return validate.Init(
		&supplyListValidator{params: p},
	).Run()
}
