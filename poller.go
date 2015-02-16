package main

type PollDesc struct {
	Schedule string
	PollType string
	Location string
	Params   []string
	DestName string
}

type Poller interface {
	poll()
}

//ipmitool -H node121.ipmi.cluster -U username -P password -I lanplus sensor
