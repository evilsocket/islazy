package main

import (
	"fmt"

	"github.com/evilsocket/islazy/zip"
)

func main() {
	if files, err := zip.Unzip("archive.zip", "./dest"); err != nil {
		panic(err)
	} else {
		fmt.Println(files)
	}
}
