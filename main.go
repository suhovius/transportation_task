package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"bitbucket.org/suhovius/transportation_task/app/forms/taskform"
	"bitbucket.org/suhovius/transportation_task/app/operations/algorithm/solver"
	"bitbucket.org/suhovius/transportation_task/app/operations/creators/taskcreator"
	"bitbucket.org/suhovius/transportation_task/app/operations/printers/taskprinter"
	"bitbucket.org/suhovius/transportation_task/app/views/errdataview"
	log "github.com/sirupsen/logrus"
)

var logFile = os.Stdout

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(logFile)
}

// TODO: Refactor handlers into separate files
func main() {
	var (
		addr = flag.String("addr", ":8080", "address of the http server")
	)

	s := NewServer(*addr)
	log.Infof("Starting server at port %s", *addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server: %v", err)
	}

}

// NewServer prepares http server.
func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	h := TaskSolvingHandler{}

	mux.Handle("/api/tasks/", &h)

	s := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return &s
}

// RequestLogger creates request logger
func RequestLogger(request *http.Request) *log.Entry {
	logger := log.WithFields(
		log.Fields{
			"method": request.Method, "url": request.URL, "ip": request.RemoteAddr,
		},
	)
	return logger
}

// TaskSolvingHandler for task solving requests
type TaskSolvingHandler struct{}

// ServerHTTP implements http.Handler.
func (h *TaskSolvingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// needs recover from panic to prevent server process exit

	logger := RequestLogger(r)

	if r.Method == "POST" {
		var err error
		var params taskform.Params

		// we can create global decoder that can decode to any structure
		// probably need to check this
		if err = json.NewDecoder(r.Body).Decode(&params); err != nil {
			message := fmt.Sprintf("JSON Decoder: %s", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusBadRequest)
			return
		}

		jsonBlob, err := json.Marshal(params)
		if err != nil {
			message := fmt.Sprintf("Marshal error: %s", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusInternalServerError)
			return
		}

		logger.Infof("Received parameters: %s", string(jsonBlob))

		// ========= Parameters Validation =====================================
		// TODO: Validate parameters cost table dimensions and supply demand list dimensions
		// TODO: Validate parameters. At least one supply and at least one demand
		// respond with http.StatusUnprocessableEntity

		// ========= Create Task Struct ========================================
		task := taskcreator.New(&params).Perform()
		logger.Info(fmt.Sprintf("Created Task UUID: %s", task.UUID))

		// TODO: Use logger interface or just our logger at taskprinter
		// instead of direct printing to output file
		// or print these table details into separate file named by task uuid
		// or smth like this
		taskprinter.New(&task, logFile).Perform()

		// ========= Find the solution =========================================
		logger.Info(fmt.Sprintf("Process Task UUID: %s", task.UUID))
		// TODO: secondsLimit might be configurable from the API
		err = solver.New(&task, 1*time.Minute).Perform()
		if err != nil {
			message := fmt.Sprintf("Task Solver: %v", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusInternalServerError)
			return
		}

		taskJSON, err := task.MarshalJSON()
		if err != nil {
			message := fmt.Sprintf("Response Rendering: %v", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusInternalServerError)
			return
		}

		w.Write(taskJSON)
	} else {
		message := "Invalid request method"
		logger.Warn(message)
		http.Error(w, APIErrorMessage(logger, message), http.StatusMethodNotAllowed)
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
