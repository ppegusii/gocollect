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
	var pollers []Poller = config.getPollers()
	for _, poller := range pollers {
		//TODO add pollers to pollerQ ordered by scheduleToTime
		s.pollerQ.Push(poller, scheduleToTime(poller.GetSchedule()))
	}
	return s
}

func (s *Scheduler) Start() {
	fmt.Printf("doing\n")
}

func scheduleToTime(schedule string) int {
	//TODO convert schedule to time since epoch
	return 0
}
