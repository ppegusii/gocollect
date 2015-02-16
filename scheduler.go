package main

import (
	"fmt"
	"github.com/oleiade/lane"
)

type Scheduler struct {
	pollers *lane.PQueue
}

func NewScheduler() *Scheduler {
	s := new(Scheduler)
	s.pollers = lane.NewPQueue(lane.MINPQ)
	return s
}

func (s *Scheduler) Start() {
	fmt.Printf("doing\n")
}
