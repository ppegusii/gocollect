package main

type ActionDesc struct {
	ActionType string
	Arguments  []string
	Name       string
	Params     map[string]string
}

type Action interface {
	do(map[string]string)
}
