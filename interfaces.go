package main

// AlgorithmStep is an interface that should be implemented by
// each service object that defines step of the transportation task solving
// algorithm
type AlgorithmStep interface {
	// TODO maybe this description should be moved to some different serivice object
	// like step printer. Each step should have it's own printer
	Description() string
	Perform() error
	ResultMessage() string
}
