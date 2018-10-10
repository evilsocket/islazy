package main

import (
	"fmt"

	"github.com/evilsocket/islazy/plugin"
)

var defines = map[string]interface{}{
	"log": func(s string) interface{} {
		fmt.Println(s)
		return nil
	},
}

func main() {
	plug, err := plugin.Load("examples/plugin.js", defines)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	ret, err := plug.Call("Hello")
	if err != nil {
		fmt.Printf("error while calling Hello function: %v\n", err)
	} else if ret != nil {
		fmt.Printf("Hello returned %v\n", ret)
	}
}
