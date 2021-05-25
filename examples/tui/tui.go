package main

import (
	"fmt"
	"os"

	"github.com/evilsocket/islazy/tui"
)

func main() {
	if !tui.Effects() {
		fmt.Printf("tui effects not available on this terminal.\n")
		return
	}

	// ASCII tables on the terminal
	columns := []string{
		"City",
		"Temp",
	}

	rows := [][]string{
		{"Roma", "16 C"},
		{"Milano", "17 C"},
		{"Firenze", "15 C"},
	}

	tui.Table(os.Stdout, columns, rows)

	err := fmt.Errorf("example error")
	// and colors :)
	if err != nil {
		fmt.Printf("\nERROR: %s\n", tui.Bold(tui.Red(err.Error())))
	} else {
		fmt.Printf("\n%s\n", tui.Green("Ok"))
	}
}
