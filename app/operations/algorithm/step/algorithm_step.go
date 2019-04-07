package step

// AlgorithmStep is an interface that should be implemented by
// each service object that defines step of the transportation task solving
// algorithm
type AlgorithmStep interface {
	Description() string
	Perform() error
	ResultMessage() string
}
