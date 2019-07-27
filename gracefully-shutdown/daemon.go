package main

import (
	"log"
	"os"
	"sync"
)

// DaemonChannel is channel to receive terminate commands
var DaemonChannel chan os.Signal

// DaemonOn represent whether use WaitGroup
var DaemonOn bool

// DaemonWaitGroup counts remaining works
var DaemonWaitGroup sync.WaitGroup

// GracefullyShutdown waits for all slaves to finish works on-hand and shutdown
func GracefullyShutdown() {
	signal := <-DaemonChannel
	numberWait := 3
	log.Printf("caught signal: %+v\n", signal)
	log.Printf("wait for %d slaves' work finished...\n", numberWait)
	DaemonWaitGroup.Add(numberWait)
	DaemonOn = true
	DaemonWaitGroup.Wait()
	os.Exit(0)
}
