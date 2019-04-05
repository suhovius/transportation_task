package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
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

	logger := RequestLogger(r)

	if r.Method == "POST" {
		var err error
		var params TaskParams

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

		task := (&TaskCreator{params: &params}).Perform()
		logger.Info(fmt.Sprintf("Created Task UUID: %s", task.UUID))
		// TODO: Refactor task printer into service object
		task.Print()

		// ========= Find the solution =========================================
		logger.Info(fmt.Sprintf("Process Task UUID: %s", task.UUID))
		// TODO: secondsLimit might be configurable from the API
		err = (&TaskSolver{task: &task, secondsLimit: 1 * time.Minute}).Peform()
		if err != nil {
			message := fmt.Sprintf("Task Solver: %v", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusInternalServerError)
			return
		}

		taskJSON, err := json.Marshal(task)
		if err != nil {
			message := fmt.Sprintf("Response Rendering: %v", err)
			http.Error(w, APIErrorMessage(logger, message), http.StatusInternalServerError)
			return
		}

		w.Write(taskJSON)
		// TODO: Round numners in api response generation and return int values there
		// https://yourbasic.org/golang/round-float-to-int/
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

	jsonBlob, err := json.Marshal(ErrorData{Message: message})
	if err != nil {
		logger.Warnf("Marshal error: %s", err)
	}
	return string(jsonBlob)
}
