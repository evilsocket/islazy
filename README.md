<p align="center">
  <p align="center">
    <a href="https://github.com/evilsocket/islazy/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/evilsocket/islazy.svg?style=flat-square"></a>
    <a href="https://github.com/evilsocket/islazy/blob/master/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-GPL3-brightgreen.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/evilsocket/islazy"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/evilsocket/islazy?style=flat-square&fuckgithubcache=1"></a>
    <a href="http://godoc.org/github.com/evilsocket/islazy">
        <img alt="Docs" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square">
    </a>
  </p>
</p>

evilsocket/islazy is a Go package containing a set of objects, helpers and functions I often find myself using in my projects.

    go get -u github.com/evilsocket/islazy

## Examples

*[islazy/tui](https://godoc.org/github.com/evilsocket/islazy/tui)*

```go
// ASCII tables on the terminal
columns := []string {
    "City",
    "Temp",
}

rows := [][]string {
    { "Roma", "16 C" },
    { "Milano", "17 C" },
    { "Firenze", "15 C" },
}

tui.Table(os.Stdout, columns, rows)

// and colors :)
if err != nil {
    fmt.Printf("ERROR: %s\n", tui.Bold(tui.Red(err.Error())) )
} else {
    fmt.Printf("%s\n", tui.Green("Ok"))
}
```

## License

This library was made with ♥  by [Simone Margaritelli](https://www.evilsocket.net/) and it's released under the GPL 3 license.
