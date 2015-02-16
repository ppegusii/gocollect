package main

type DataDestDesc struct {
	Name     string
	DestType string
	Location string
	Params   []string
}

type DataDest interface {
	write(string)
}
