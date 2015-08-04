package main

import (
	"fmt"
	"github.com/oleiade/lane"
	"log"
	"math"
	"strconv"
	"time"
)

type Scheduler struct {
	pollerQ *lane.PQueue
}

func NewScheduler(configFileName *string) *Scheduler {
	s := new(Scheduler)
	s.pollerQ = lane.NewPQueue(lane.MINPQ)
	var config *Config = parse(configFileName)
	var pollers []Poller = config.getPollers()
	log.Printf("pollers = %+v", pollers)
	for _, poller := range pollers {
		//TODO add pollers to pollerQ ordered by scheduleToTime
		log.Printf("Pushing %+v", poller)
		s.pollerQ.Push(poller, scheduleToTime(poller.GetSchedule()))
	}
	//log.Printf(s.pollerQ.Size())
	return s
}

func (s *Scheduler) Start() {
	fmt.Printf("doing\n")
	// TODO:
	// While polling job at the head of the Q should be executed:
	// Remove polling job from the head of the Q.
	//	Execute polling job in new routine.
	//	Place polling job back on Q with the next run time as the priority.
	// Sleep until the next scheduled poll should execute.
	for {
		var curTime, nextTime, sleepTime int
		var nextPoller Poller
		var next interface{}
		curTime = int(time.Now().Unix())
		next, nextTime = s.pollerQ.Pop()
		nextPoller, _ = next.(Poller)
		sleepTime = nextTime - curTime
		log.Printf("Sleeping for %d seconds...", sleepTime)
		if sleepTime > 0 {
			time.Sleep(time.Second * time.Duration(sleepTime))
		}
		log.Printf("polling with poller: %+v", nextPoller)
		go nextPoller.Poll()
		s.pollerQ.Push(nextPoller, scheduleToTime(nextPoller.GetSchedule()))
	}
}

func scheduleToTime(schedule string) int {
	//TODO add ability to parse cron like strings
	var i int
	var err error
	i, err = strconv.Atoi(schedule)
	if err != nil {
		log.Fatal(err)
	}
	//Convert milliseconds to seconds
	i = int(math.Ceil(float64(i) / 1E3))
	return int(time.Now().Unix()) + i
}
