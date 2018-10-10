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

**[islazy/async](https://godoc.org/github.com/evilsocket/islazy/async)**

```go
// timed calls
ret, err := async.WithTimeout( 5 * time.Second, func() interface{}{
    // do something here
    return nil
})

// multi core work queue
func logic(arg async.Job) {
    fmt.Printf("got job %d", arg.(int))
}

q := async.NewQueue(0, logic)

for i := 0; i < 10; i++ {
    q.Add(async.Job(i))
}

q.WaitDone()
```

**[islazy/plugin](https://godoc.org/github.com/evilsocket/islazy/plugin)**

From [the shellz project](https://github.com/evilsocket/shellz):

```go
type Plugin struct {
	*plugin.Plugin

	timeouts core.Timeouts
	ctx      interface{}
}

func LoadPlugin(path string) (error, *Plugin) {
	if p, err := plugin.Load(path, defines); err != nil {
		return err, nil
	} else {
		return nil, &Plugin{
			Plugin: p,
		}
	}
}

func (p *Plugin) NewSession(sh models.Shell, timeouts core.Timeouts) (err error, clone *Plugin) {
	p.Lock()
	defer p.Unlock()

	if err, clone = LoadPlugin(p.Path); err != nil {
		return
	}
    // ... cut ...
    clone.ctx, err = clone.Call("Create", sh)
    // ... cut ...
	return
}

// ... cut ...

func (p *Plugin) Exec(cmd string) ([]byte, error) {
	p.Lock()
	defer p.Unlock()
// ... cut ...
    if ret, err := p.Call("Exec", p.ctx, cmd); err != nil {
        return eres{err: err}
    } 
// ... cut ...
}

func (p *Plugin) Close() {
	p.Lock()
	defer p.Unlock()

	if err, _ := p.Call("Close", p.ctx); err != nil {
		log.Warning("error while running Close callback for plugin %s: %s", p.Name, err)
	}
}
```

## License

This library was made with ♥  by [Simone Margaritelli](https://www.evilsocket.net/) and it's released under the GPL 3 license.
