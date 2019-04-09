package solver

import (
	"fmt"
	"time"

	"bitbucket.org/suhovius/transportation_task/app/models/taskmodel"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/sequence"
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
	log "github.com/sirupsen/logrus"
)

// TaskSolver provides transport task solution finding algorithm logic
type TaskSolver struct {
	task         *taskmodel.Task
	secondsLimit time.Duration
	startTime    time.Time
	elapsedTime  time.Duration
	logEntry     *log.Entry
}

// New returns new TaskSolver instance
func New(
	task *taskmodel.Task, secondsLimit time.Duration, logEntry *log.Entry,
) *TaskSolver {
	return &TaskSolver{task: task, secondsLimit: secondsLimit, logEntry: logEntry}
}

// Perform finds transport task solution
func (ts *TaskSolver) Perform() (err error) {
	ts.startTime = time.Now()

	ts.logEntry.Info("=== Initial Preparations ===")
	err = ts.runInitialSteps()

	if err != nil {
		return err
	}

	ts.printSolutionPrice()

	iterationsCount, err := ts.runPotentialsMethod()

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	ts.logEntry.Infof(
		"=== Caclulation took %s and %d iterations ===",
		ts.elapsedTime, iterationsCount,
	)

	return
}

func (ts *TaskSolver) printSolutionPrice() {
	ts.logEntry.Infof(
		"Delivery Cost: %d",
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

func (ts *TaskSolver) runInitialSteps() error {
	return sequence.New(
		balance.New(ts.task),
		degeneracyprev.New(ts.task),
		// TODO: Add Minimal Rates and Vogel approximation methods
		// Add ability to select approximation method via API
		northwestcrnr.New(ts.task),
	).RunWithLog(ts.logEntry)
}

func (ts *TaskSolver) runIterationSteps() error {
	return sequence.New(
		iterationinit.New(ts.task),
		amountdistribcheck.New(ts.task),
		degeneracycheck.New(ts.task),
		potentialcalc.New(ts.task),
		optsolcheck.New(ts.task),
		circuitbuild.New(ts.task),
		supplyredistrib.New(ts.task),
	).RunWithLog(ts.logEntry)
}

func (ts *TaskSolver) runPotentialsMethod() (iterationsCount int, err error) {
	for i := 1; !ts.task.IsOptimalSolution; i++ {
		iterationsCount = i
		ts.logEntry.Infof("=== Potentials Method. Iteration #%d ===", i)
		err = ts.checkTimeLimit()

		if err != nil {
			break
		}

		err = ts.runIterationSteps()

		if err != nil {
			break
		}
	}

	return
}
