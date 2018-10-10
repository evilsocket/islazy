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

**[islazy/tui](https://godoc.org/github.com/evilsocket/islazy/tui)**

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

**[islazy/str](https://godoc.org/github.com/evilsocket/islazy/str)**

```go
// trimming
str.Trim(" hello world ") == "hello world"
str.TrimRight(" hello world ") == " hello world"
str.TrimLeft(" hello world ") == "hello world "

// splitting
str.SplitBy("1,2,,3,4,", "," ) == []string{ "1", "2", "3", "4" }
str.Comma("1,2,,3,4,") == []string{ "1", "2", "3", "4" }
```

**[islazy/fs](https://godoc.org/github.com/evilsocket/islazy/fs)**

```go
// misc
fs.Expand("~/") -> "/home/evilsocket"
fs.Exists("nope") == false

// globbing
fs.Glob( "/some/path/", "*.log", cn fun(fileName string) error {
    // do something with fileName
    return nil
})

// reading
lines, err := fs.LineReader("/some/file.log")
for line := range lines {

}
```

## License

This library was made with ♥  by [Simone Margaritelli](https://www.evilsocket.net/) and it's released under the GPL 3 license.
