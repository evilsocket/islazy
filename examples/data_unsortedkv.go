package main

import (
	"fmt"

	"github.com/evilsocket/islazy/data"
)

func main() {
	// create a new unsorted key value store on /tmp/ukv.dat
	// it'll be flushed on each modification
	kv, err := data.NewDiskUnsortedKV("/tmp/ukv.dat")
	if err != nil {
		panic(err)
	}
	// set a few values
	kv.Set("foo", "bar")
	kv.Set("moo", "boo")
	kv.Set("burn", "the witches")

	// load it
	kv2, err := data.NewDiskUnsortedKVReader("/tmp/ukv.dat")
	if err != nil {
		panic(err)
	}
	// print each object (they'll be in a different order each time)
	kv2.Each(func(key, value string) bool {
		fmt.Printf("%s => %s\n", key, value)
		// don't stop the loop
		return false
	})
}
