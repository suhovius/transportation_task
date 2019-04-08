package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"bitbucket.org/suhovius/transportation_task/app/actions/solvetaskhandler"
	"bitbucket.org/suhovius/transportation_task/utils/requestid"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

// TODO needs recover from panic to prevent server process exit

// Add middlewares for tracing, logging and recovery

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
	router := http.NewServeMux()

	logger := log.New()

	router.Handle("/api/tasks/", solvetaskhandler.New(logger))

	nextRequestID := requestid.Next

	s := http.Server{
		Addr:    addr,
		Handler: tracing(nextRequestID)(logging(logger)(router)),
	}

	return &s
}

func logLine(logger *log.Logger, word string, r *http.Request) {
	requestID, ok := r.Context().Value(requestid.RequestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}
	logger.Println(word, requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logLine(logger, "Started", r)
			defer func() {
				logLine(logger, "Finished", r)
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestid.RequestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
