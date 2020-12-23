package main

import (
	"fmt"

	"github.com/evilsocket/islazy/fs"
)

func main() {
	// misc
	fmt.Println(fs.Expand("~/"))
	fmt.Println(fs.Exists("nope"))

	// reading
	if lines, err := fs.LineReader("/etc/passwd"); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		for line := range lines {
			fmt.Println(line)
		}
	}

	// change working directory
	err := fs.Chdir("/", func() error {
		// glob
		return fs.Glob(".", "*.*", func(fileName string) error {
			fmt.Printf("%s\n", fileName)
			return nil
		})
	})
	if err != nil {
		panic(err)
	}
}
