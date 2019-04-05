package main

import (
	"fmt"
	"time"
)

// TaskSolver provides transport task solution finding algorithm logic
type TaskSolver struct {
	task         *Task
	secondsLimit time.Duration
	startTime    time.Time
	elapsedTime  time.Duration
}

// Peform finds transport task solution
func (ts *TaskSolver) Peform() (err error) {
	ts.startTime = time.Now()

	fmt.Printf("\n=== Initial Preparations =================================\n")
	err = ts.createInitialSequence().Run()

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	for i := 1; !ts.task.IsOptimalSolution; i++ {
		fmt.Printf("\n=== Potentials Method. Iteration #%d ==============\n", i)
		err = ts.checkTimeLimit()

		if err != nil {
			break
		}

		err = ts.createIterativeSequence().Run()

		if err != nil {
			break
		}
	}

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	fmt.Printf("Caclulation took %s\n", ts.elapsedTime)

	return
}

func (ts *TaskSolver) createInitialSequence() *StepsSequencePerformer {
	var initialSteps []AlgorithmStep
	initialSteps = append(
		initialSteps,
		&Balancer{task: ts.task},
		&DegeneracyPreventer{task: ts.task},
		&NorthWestCornerSolutionFinder{task: ts.task},
	)
	return &StepsSequencePerformer{task: ts.task, steps: &initialSteps}
}

func (ts *TaskSolver) createIterativeSequence() *StepsSequencePerformer {
	var iterativeSteps []AlgorithmStep
	iterativeSteps = append(
		iterativeSteps,
		&IterationInitializer{task: ts.task},
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

func (ts *TaskSolver) checkTimeLimit() (err error) {
	ts.elapsedTime = time.Since(ts.startTime)
	if ts.elapsedTime > ts.secondsLimit {
		err = fmt.Errorf(
			"Calculation took %s and exceded allowed limit of %s",
			ts.elapsedTime, ts.secondsLimit,
		)
	}
	return
}

// Add some kind of printer object or output stream to send all the print
// requests and to be able to change where they are sent to
