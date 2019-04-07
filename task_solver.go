package main

import (
	"fmt"
	"time"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/step"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/amountdistribcheck"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/balance"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/circuitbuild"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/degeneracycheck"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/degeneracyprev"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/iterationinit"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/northwestcrnr"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/optsolcheck"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/potentialcalc"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/steps/supplyredistrib"
	"bitbucket.org/suhovius/transportation_task/app/operations/deliverycost"
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

func (ts *TaskSolver) newSequencePerformer(steps ...step.AlgorithmStep) *StepsSequencePerformer {
	return &StepsSequencePerformer{task: ts.task, steps: &steps}
}

func (ts *TaskSolver) createInitialSequence() *StepsSequencePerformer {
	return ts.newSequencePerformer(
		balance.New(ts.task),
		degeneracyprev.New(ts.task),
		northwestcrnr.New(ts.task),
	)
}

func (ts *TaskSolver) createIterativeSequence() *StepsSequencePerformer {
	return ts.newSequencePerformer(
		iterationinit.New(ts.task),
		amountdistribcheck.New(ts.task),
		degeneracycheck.New(ts.task),
		potentialcalc.New(ts.task),
		optsolcheck.New(ts.task),
		circuitbuild.New(ts.task),
		supplyredistrib.New(ts.task),
	)
}

func (ts *TaskSolver) printSolutionPrice() {
	fmt.Printf(
		"Delivery Cost: %d\n",
		mathext.RoundToInt(
			deliverycost.New(ts.task).Perform(),
		),
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
