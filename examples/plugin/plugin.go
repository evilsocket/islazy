package main

import (
	"fmt"

	"github.com/evilsocket/islazy/plugin"
)

func main() {
	// define some functions
	plugin.Defines = map[string]interface{}{
		"log": func(s string) interface{} {
			fmt.Println(s)
			return nil
		},
	}

	// load the plugin
	plug, err := plugin.Load("examples/plugin.js")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// call methods
	methods := plug.Methods()
	for _, m := range methods {
		ret, err := plug.Call(m)
		if err != nil {
			fmt.Printf("error while calling %s function: %v\n", m, err)
		} else if ret != nil {
			fmt.Printf("plugin.%s -> %v\n", m, ret)
		}

	}

	// get variables
	obj, err := plug.GetObject("Text")
	fmt.Printf("plugin.Text = '%s'\n", obj.(string))
}
