package main

import (
	"fmt"

	"github.com/evilsocket/islazy/ops"
)

func main() {
	s := ops.Ternary(1 == 2, "equal", "not equal").(string)
	fmt.Println(s)
	s = ops.Ternary(2 == 2, "equal", "not equal").(string)
	fmt.Println(s)
}
