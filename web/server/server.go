package server

import (
	"context"
	"net/http"

	"bitbucket.org/suhovius/transportation_task/app/actions/solvetaskhandler"
	"bitbucket.org/suhovius/transportation_task/utils/requestid"
	log "github.com/sirupsen/logrus"
)

// TODO needs recover from panic to prevent server process exit

// New prepares http server.
func New(addr string, logger *log.Logger) *http.Server {
	router := http.NewServeMux()

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

// logging middleware
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

// tracing middleware
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
