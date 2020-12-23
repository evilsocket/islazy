package main

import (
	"fmt"

	"github.com/evilsocket/islazy/str"
)

func main() {
	s := "\t hello world "
	csv := "1,2,, ,3,4,"

	fmt.Println(str.Trim(s))
	fmt.Println(str.TrimRight(s))
	fmt.Println(str.TrimLeft(s))
	fmt.Println(str.SplitBy(csv, ","))
	fmt.Println(str.Comma(csv))
}
