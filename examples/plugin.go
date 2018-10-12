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

	methods := plug.Methods()
	for _, m := range methods {
		ret, err := plug.Call(m)
		if err != nil {
			fmt.Printf("error while calling %s function: %v\n", m, err)
		} else if ret != nil {
			fmt.Printf("%s returned %v\n", m, ret)
		}

	}

	obj, err := plug.GetObject("Text")
	fmt.Printf("plugin.Text = '%s'\n", obj.(string))

}
