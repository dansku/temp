package main

import (
	"log"

	"github.com/dansku/metering/scheduler"
)

func main() {

	log.Printf("Starting Application")

	// Start Scheduler
	scheduler.Start()

}
