package scheduler

import (
	"log"
	"time"
)

// Start scheduler
func Start() {
	log.Printf("Starting Scheduler!")

	// Hack to run forever
	i := false
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Scheduler running...")
		if i == true {
			break
		}
	}
}
