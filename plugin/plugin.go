package plugin

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/robertkrimen/otto"
)

// Plugin is an object representing a javascript
// file exporting functions and variables that
// your project can use to extend its functionalities.
type Plugin struct {
	sync.Mutex
	// The basename of the plugin.
	Name string
	// The actual javascript code.
	Code string
	// The full path of the plugin.
	Path string

	vm        *otto.Otto
	callbacks map[string]otto.Value
	objects   map[string]otto.Value
}

// Load loads and compiles a plugin given its path.
func Load(path string) (error, *Plugin) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err, nil
	}

	plugin := &Plugin{
		Name:      strings.Replace(filepath.Base(path), ".js", "", -1),
		Code:      string(raw),
		Path:      path,
		callbacks: make(map[string]otto.Value),
		objects:   make(map[string]otto.Value),
	}

	if err = plugin.compile(); err != nil {
		return err, nil
	}

	return nil, plugin
}

// Clone returns a new instance identical to the plugin.
func (p *Plugin) Clone() *Plugin {
	_, clone := Load(p.Path)
	return clone
}

// Call executes one of the declared callbacks of the plugin by its name.
func (p *Plugin) Call(name string, args ...interface{}) (error, interface{}) {
	if cb, found := p.callbacks[name]; !found {
		return fmt.Errorf("%s does not name a function", name), nil
	} else if ret, err := cb.Call(otto.NullValue(), args...); err != nil {
		return err, nil
	} else if !ret.IsUndefined() {
		if exported, err := ret.Export(); err != nil {
			return err, nil
		} else {
			return nil, exported
		}
	}
	return nil, nil
}
