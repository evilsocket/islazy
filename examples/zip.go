package main

import (
	"fmt"

	"github.com/evilsocket/islazy/zip"
)

func main() {
	if err := zip.Files("archive.zip", []string{"README.md", "release.sh"}); err != nil {
		panic(err)
	}

	if files, err := zip.Unzip("archive.zip", "./dest"); err != nil {
		panic(err)
	} else {
		fmt.Println(files)
	}
}
