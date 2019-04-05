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

const (
	apiPathPrefix = "api"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	var (
		addr = flag.String("addr", ":8080", "address of the http server")
	)

	s := NewServer(*addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server: %v", err)
	}
}

// NewServer prepares http server.
func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	h := TaskSolvingHandler{}

	mux.Handle(apiPathPrefix+"/tasks", &h)

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
			logger.Fatalf("JSON Decoder: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		logger.Info(fmt.Sprintf("Received parameters: %v", r.Body))

		// ========= Parameters Validation =====================================
		// TODO: Validate parameters cost table dimensions and supply demand list dimensions
		// TODO: Validate parameters. At least one supply and at least one demand

		// ========= Create Task Struct ========================================

		task := (&TaskCreator{params: &params}).Perform()
		logger.Info(fmt.Sprintf("Created Task UUID: %s", task.UUID))
		// TODO: Refactor task printer into service object
		task.Print()

		// ========= Find the solution =========================================
		// TODO: secondsLimit might be configurable from the API
		err = (&TaskSolver{task: &task, secondsLimit: 10 * time.Minute}).Peform()
		if err != nil {
			message := fmt.Sprintf("Task Solver: %v", err)
			logger.Fatal(message)
			http.Error(w, message, http.StatusInternalServerError)
		}

		taskJSON, err := json.Marshal(task)
		if err != nil {
			message := fmt.Sprintf("Response Rendering: %v", err)
			logger.Fatal(message)
			http.Error(w, message, http.StatusInternalServerError)
		}

		w.Write(taskJSON)
		// TODO: Round numners in api response generation and return int values there
		// https://yourbasic.org/golang/round-float-to-int/
	} else {
		message := "Invalid request method"
		logger.Fatal(message)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}
