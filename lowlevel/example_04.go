package main

import "fmt"

//go:generate stringer -type=State
type State int

const (
	Created State = iota
	Starting
	Started
	Done
)

func Example04() {

	fmt.Println(Created, Starting, Started, Done)
}
