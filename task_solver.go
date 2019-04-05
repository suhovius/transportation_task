package main

import "fmt"

// TaskSolver provides transport task solution finding algorithm logic
type TaskSolver struct {
	task *Task
}

// Peform finds transport task solution
func (ts *TaskSolver) Peform() (err error) {
	err = ts.defineInitialLoop().Run()

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	err = ts.defineIterativeLoop().Run()

	if err != nil {
		return
	}

	return
}

func (ts *TaskSolver) defineInitialLoop() *StepsSequencePerformer {
	var initialSteps []AlgorithmStep
	initialSteps = append(
		initialSteps,
		&Balancer{task: ts.task},
		&DegeneracyPreventer{task: ts.task},
		&NorthWestCornerSolutionFinder{task: ts.task},
	)
	return &StepsSequencePerformer{task: ts.task, steps: &initialSteps}
}

func (ts *TaskSolver) defineIterativeLoop() *StepsSequencePerformer {
	var iterativeSteps []AlgorithmStep
	iterativeSteps = append(
		iterativeSteps,
		&AmountDistributionChecker{task: ts.task},
		&DegeneracyChecker{task: ts.task},
		&PotentialsCalculator{task: ts.task},
		&OptimalSolutionChecker{task: ts.task},
		&CircuitBuilder{task: ts.task},
		&SupplyRedistributor{task: ts.task},
	)

	return &StepsSequencePerformer{task: ts.task, steps: &iterativeSteps}
}

func (ts *TaskSolver) printSolutionPrice() {
	fmt.Printf("Delivery Cost: %d\n", roundToInt(
		(&DeliveryCostCalculator{task: ts.task}).Peform()),
	)
}

// TODO: Check Cycles Count limit or finding time like 1 minute for example

// TODO: Clear/Reset previous values from calculation

// Later each step could be started with step runner service object wrapper
