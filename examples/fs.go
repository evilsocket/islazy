package main

import (
	"fmt"

	"github.com/evilsocket/islazy/fs"
)

func main() {
	// misc
	fmt.Println(fs.Expand("~/"))
	fmt.Println(fs.Exists("nope"))

	// globbing
	fs.Glob(".", "*.*", func(fileName string) error {
		// do something with fileName
		return nil
	})

	// reading
	if lines, err := fs.LineReader("/some/file.log"); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		for line := range lines {
			fmt.Println(line)
		}
	}
}
