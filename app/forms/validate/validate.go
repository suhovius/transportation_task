package validate

// ValidationChain contains list of validators
type ValidationChain struct {
	validators *[]Validator
}

// Init creates new ValidationChain with provided list of validators
func Init(validators ...Validator) *ValidationChain {
	return &ValidationChain{validators: &validators}
}

// Run runs validators from ValidationChain returning first error
func (vc *ValidationChain) Run() (err error) {
	for _, validator := range *vc.validators {
		err = validator.Perform()
		if err != nil {
			break
		}
	}

	return
}
