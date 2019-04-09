package solvetaskhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"bitbucket.org/suhovius/transportation_task/app/forms/taskform"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/solver"
	"bitbucket.org/suhovius/transportation_task/app/operations/creators/taskcreator"
	"bitbucket.org/suhovius/transportation_task/app/operations/printers/taskprinter"
	"bitbucket.org/suhovius/transportation_task/app/views/errdataview"
	"bitbucket.org/suhovius/transportation_task/utils/requestid"
	log "github.com/sirupsen/logrus"
)

// TaskSolvingHandler for task solving requests
type TaskSolvingHandler struct {
	logger *log.Logger
}

// New returns new TaskSolvingHandler
func New(logger *log.Logger) *TaskSolvingHandler {
	return &TaskSolvingHandler{logger: logger}
}

func (h *TaskSolvingHandler) logEntryBy(request *http.Request) *log.Entry {
	requestID, ok := request.Context().Value(requestid.RequestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}
	return h.logger.WithFields(log.Fields{"request_id": requestID})
}

// ServerHTTP implements http.Handler.
func (h *TaskSolvingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	le := h.logEntryBy(r)

	if r.Method == "POST" {
		var err error
		var params taskform.Params

		// we can create global decoder that can decode to any structure
		// probably need to check this
		if err = json.NewDecoder(r.Body).Decode(&params); err != nil {
			message := fmt.Sprintf("JSON Decoder: %s", err)
			http.Error(w, APIErrorMessage(le, message), http.StatusBadRequest)
			return
		}

		jsonBlob, err := json.Marshal(params)
		if err != nil {
			message := fmt.Sprintf("Marshal error: %s", err)
			http.Error(w, APIErrorMessage(le, message), http.StatusInternalServerError)
			return
		}

		le.Infof("Received parameters: %s", string(jsonBlob))

		// ========= Parameters Validation =====================================
		// TODO: Validate parameters cost table dimensions and supply demand list dimensions
		// TODO: Validate parameters. At least one supply and at least one demand
		// respond with http.StatusUnprocessableEntity
		err = params.Validate()
		if err != nil {
			message := fmt.Sprintf("Params Validation Error: %s", err)
			http.Error(w, APIErrorMessage(le, message), http.StatusUnprocessableEntity)
			return
		}

		// ========= Create Task Struct ========================================
		task := taskcreator.New(&params).Perform()
		le.Info(fmt.Sprintf("Created Task UUID: %s", task.UUID))

		// TODO: Print detailed log (with ASCII tables) into separate file
		taskprinter.New(&task, os.Stdout).Perform()

		// ========= Find the solution =========================================
		le.Info(fmt.Sprintf("Process Task UUID: %s", task.UUID))
		// TODO: secondsLimit might be configurable from the API
		err = solver.New(&task, 1*time.Minute, le).Perform()
		if err != nil {
			message := fmt.Sprintf("Task Solver: %v", err)
			http.Error(w, APIErrorMessage(le, message), http.StatusInternalServerError)
			return
		}

		taskJSON, err := task.MarshalJSON()
		if err != nil {
			message := fmt.Sprintf("Response Rendering: %v", err)
			http.Error(w, APIErrorMessage(le, message), http.StatusInternalServerError)
			return
		}

		w.Write(taskJSON)
	} else {
		message := "Invalid request method"
		le.Warn(message)
		http.Error(w, APIErrorMessage(le, message), http.StatusMethodNotAllowed)
		return
	}
}

// APIErrorMessage creates ErrorData struct
func APIErrorMessage(logger *log.Entry, message string) string {
	logger.Warn(message)

	jsonBlob, err := json.Marshal(errdataview.New(message))
	if err != nil {
		logger.Warnf("Marshal error: %s", err)
	}
	return string(jsonBlob)
}
