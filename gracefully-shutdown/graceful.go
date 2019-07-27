package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	plannedWork := 10
	DaemonOn = false
	DaemonChannel = make(chan os.Signal)
	go GracefullyShutdown()
	signal.Notify(DaemonChannel, syscall.SIGTERM)
	signal.Notify(DaemonChannel, syscall.SIGINT)

	for i := 0; i < plannedWork; i++ {
		go SlavesWork(i)
		time.Sleep(1 * time.Second)
	}
}
