package main

import (
	"fmt"
	"time"

	"github.com/evilsocket/islazy/async"
)

func main() {
	// timed calls
	_, err := async.WithTimeout(1*time.Second, func() interface{} {
		// this will timeout
		time.Sleep(5 * time.Second)
		return nil
	})
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// -1/0 = use all logical CPUs
	q := async.NewQueue(0, func(arg async.Job) {
		fmt.Printf("got job %d\n", arg.(int))
	})

	for i := 0; i < 10; i++ {
		q.Add(async.Job(i))
	}

	q.WaitDone()
}
