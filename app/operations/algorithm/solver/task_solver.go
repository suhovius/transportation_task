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
	err = sequence.New(
		balance.New(ts.task),
		degeneracyprev.New(ts.task),
		northwestcrnr.New(ts.task),
	).RunWithLog(ts.logEntry)

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	// This part can also be splitted into separate method or struct. Will see later
	// with wrapper that prints solution price or this might be a config of
	// TaskPrinter service object

	iterationNum := 0
	for i := 1; !ts.task.IsOptimalSolution; i++ {
		iterationNum = i
		ts.logEntry.Infof("=== Potentials Method. Iteration #%d ===", i)
		err = ts.checkTimeLimit()

		if err != nil {
			break
		}

		err = sequence.New(
			iterationinit.New(ts.task),
			amountdistribcheck.New(ts.task),
			degeneracycheck.New(ts.task),
			potentialcalc.New(ts.task),
			optsolcheck.New(ts.task),
			circuitbuild.New(ts.task),
			supplyredistrib.New(ts.task),
		).RunWithLog(ts.logEntry)

		if err != nil {
			break
		}
	}

	if err != nil {
		return
	}

	ts.printSolutionPrice()

	// TODO: Print this to logger
	ts.logEntry.Infof("=== Caclulation took %s and %d iterations ===", ts.elapsedTime, iterationNum)

	return
}

func (ts *TaskSolver) printSolutionPrice() {
	ts.logEntry.Infof(
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
