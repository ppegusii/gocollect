package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ActionDescs     []ActionDesc
	PollerDescs     []PollerDesc
	ConnectionDescs []ConnectionDesc
}

/*
type Config struct {
	ActionDescs     []ActionDesc
	Schedule []PollerDesc
	ConnectionDescs []ConnectionDesc
}
*/

func (this *Config) getPollers() []Poller {
	var pollers []Poller = make([]Poller, len(this.PollerDescs))
	var pollerFactory = NewPollerFactory()
	log.Printf("this.PollerDescs = %+v", this.PollerDescs)
	for i, pollerDesc := range this.PollerDescs {
		log.Printf("pollerDesc = %+v", pollerDesc)
		pollers[i] = pollerFactory.Construct(&pollerDesc)
	}
	return pollers
}

func parse(configFileName *string) *Config {
	var file []byte
	var err error
	file, err = ioutil.ReadFile(*configFileName)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("JSON error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Config: %+v\n", config)
	return &config
}
