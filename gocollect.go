package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFileName *string = flag.String("c", "", "config file")
	flag.Parse()
	fmt.Printf("Starting Collector\n")
	scheduler := NewScheduler(configFileName)
	scheduler.Start() //just a place holder
}
