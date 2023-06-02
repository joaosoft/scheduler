package main

import (
	"github.com/joaosoft/scheduler"
)

func main() {
	s, err := scheduler.NewScheduler()
	if err != nil {
		panic(err)
	}

	if err = s.Start(); err != nil {
		panic(err)
	}
}
