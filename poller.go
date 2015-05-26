package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

type PollerDesc struct {
	ActionNames map[string]string
	Arguments   []string
	Name        string
	Params      map[string]string
	PollType    string
	Schedule    string
}

type Poller interface {
	GetSchedule() string
	Poll()
}

type PollerConstructor func(*PollerDesc) Poller

type PollerFactory struct {
	pollerConstructors map[string]PollerConstructor
}

func NewPollerFactory() *PollerFactory {
	return &PollerFactory{
		pollerConstructors: map[string]PollerConstructor{
			"Bash": NewBash,
		},
	}
}

func NewBash(pd *PollerDesc) Poller {
	return &Bash{
		command: pd.Params["command"],
		Name:    pd.Name,
		//TODO onSuccess
		schedule: pd.Schedule,
	}
}

type Bash struct {
	arguments []string
	command   string
	Name      string
	onSuccess *Action
	schedule  string
}

func (b *Bash) Poll() {
	var cmd *exec.Cmd = exec.Command(b.command, b.arguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("stdout of %s %v is %s\n", b.command, b.arguments, out.String())
}

func (b *Bash) GetSchedule() string {
	return b.schedule
}

//ipmitool -H node121.ipmi.cluster -U username -P password -I lanplus sensor
