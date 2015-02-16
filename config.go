package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	DataDestDescs []DataDestDesc
	PollDescs     []PollDesc
}

//just a place holder
func (c *Config) getDests() []DataDestDesc {
	return c.DataDestDescs
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
