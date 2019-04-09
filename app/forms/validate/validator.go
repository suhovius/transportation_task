package validate

// Validator is an interface that should be implemented by all the validators
type Validator interface {
	Perform() error
}
