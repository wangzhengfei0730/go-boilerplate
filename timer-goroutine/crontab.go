package main

import (
	"log"
	"os"
	"time"
)

// check the existence of "test.finished"
func checkStatue(valve chan bool) bool {
	log.Println("check whether file exists!")
	if _, err := os.Stat("test.finished"); err == nil {
		return true
	}
	return false
}

func timerCheck(valve chan bool) {
	log.Println("goroutine is running!")
	ch := make(chan bool, 1)
	ch <- true
	shortTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case i := <-ch:
			log.Println(i)
			log.Println("5 seconds pass, let's check once more!")
			if fileExist := checkStatue(valve); fileExist == true {
				valve <- true
				return
			}
		case <-shortTicker.C:
			log.Println("internal timer")
			ch <- true
		}
	}
}

func main() {
	valve := make(chan bool)
	ticker := time.NewTicker(20 * time.Second)
	go timerCheck(valve)
	for {
		select {
		case <-valve:
			log.Println("you got it!")
		case <-ticker.C:
			log.Println("I'm waiting every 20 seconds...")
		}
	}
}
