package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nchern/homevscorona/backend/api/pkg/cli"
	"github.com/nchern/homevscorona/backend/api/pkg/srv"
)

const (
	svcName = "api"

	shortHelp = "%s is a core api service for wirvscorona hackathon"
)

var ()

func init() {
	cli.Init(
		svcName,
		"Long help",
		fmt.Sprintf(shortHelp, svcName),
		runAndWait,
	)
}

func runAndWait() {
	go srv.Start(svcName)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for range c {
		srv.Stop()
		break
	}
}

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
