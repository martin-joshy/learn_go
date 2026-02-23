package main

import "fmt"

type SeverState int

const (
	IdleState SeverState = iota
	ConnectedState
	ErrorState
	RetryingState
)

var mapStateReadable = map[SeverState]string{
	IdleState:      "idle",
	ConnectedState: "connected",
	ErrorState:     "error",
	RetryingState:  "retrying",
}

func (ss SeverState) String() string {
	return mapStateReadable[ss]
}

func transition(ss SeverState) SeverState {
	switch ss {
	case IdleState:
		return ConnectedState
	case ConnectedState, RetryingState:
		return IdleState
	case ErrorState:
		return ErrorState
	default:
		panic(fmt.Errorf("Unknow State: %s", ss))
	}
}

func main() {
	res := transition(IdleState)
	fmt.Println(res)

	res2 := transition(RetryingState)
	fmt.Println(res2)
}
