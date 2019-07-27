package main

import (
	"fmt"
)

// SlavesWork do something with goroutine
func SlavesWork(slaveID int) {
	fmt.Printf("slave id = %d\n", slaveID)
	if DaemonOn == true {
		DaemonWaitGroup.Done()
	}
}
