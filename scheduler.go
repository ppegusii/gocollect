package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

type Scheduler struct {
	pollerQ *lane.PQueue
}

func NewScheduler(configFileName *string) *Scheduler {
	s := new(Scheduler)
	s.pollerQ = lane.NewPQueue(lane.MINPQ)
	var config *Config = parse(configFileName)
	config.getDests() //just a place holder
	return s
}

func (s *Scheduler) Start() {
	fmt.Printf("doing\n")
}
