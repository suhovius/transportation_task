package main

import (
	"fmt"
	"time"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/utils/mathext"
)

// TaskSolver provides transport task solution finding algorithm logic
type TaskSolver struct {
	task         *taskmodel.Task
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

	// This part can also be splitted into separate method or struct. Will see later
	// with wrapper that prints solution price or this might be a config of
	// TaskPrinter service object

	// TODO: Return iterations number in the log
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

	// here we can add step that rounds delivery prices to int values
	// or maybe that will require additional struct with int values for the
	// response. Maybe json modificator at stuct has option to format response somehow. Will see

	ts.printSolutionPrice()

	// TODO: Print this to logger
	fmt.Printf("Caclulation took %s\n", ts.elapsedTime)

	return
}

func (ts *TaskSolver) newSequencePerformer(steps ...AlgorithmStep) *StepsSequencePerformer {
	return &StepsSequencePerformer{task: ts.task, steps: &steps}
}

func (ts *TaskSolver) createInitialSequence() *StepsSequencePerformer {
	return ts.newSequencePerformer(
		&Balancer{task: ts.task},
		&DegeneracyPreventer{task: ts.task},
		&NorthWestCornerSolutionFinder{task: ts.task},
	)
}

func (ts *TaskSolver) createIterativeSequence() *StepsSequencePerformer {
	return ts.newSequencePerformer(
		&IterationInitializer{task: ts.task},
		&AmountDistributionChecker{task: ts.task},
		&DegeneracyChecker{task: ts.task},
		&PotentialsCalculator{task: ts.task},
		&OptimalSolutionChecker{task: ts.task},
		&CircuitBuilder{task: ts.task},
		&SupplyRedistributor{task: ts.task},
	)
}

func (ts *TaskSolver) printSolutionPrice() {
	fmt.Printf("Delivery Cost: %d\n", mathext.RoundToInt(
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
