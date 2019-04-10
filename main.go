package main

import (
	"flag"
	"os"

	"bitbucket.org/suhovius/transportation_task/web/server"
	log "github.com/sirupsen/logrus"
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

	s := server.New(*addr, log.New())

	log.Infof("Starting server at port %s", *addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
