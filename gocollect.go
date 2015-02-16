package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Starting Collector\n")
	scheduler := NewScheduler()
	scheduler.Start()
}
