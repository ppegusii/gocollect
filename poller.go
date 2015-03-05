package main

type PollerDesc struct {
	ActionNames map[string]string
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
	command   string
	Name      string
	onSuccess *Action
	schedule  string
}

func (b *Bash) Poll() {
}

func (b *Bash) GetSchedule() string {
	return b.schedule
}

//ipmitool -H node121.ipmi.cluster -U username -P password -I lanplus sensor
