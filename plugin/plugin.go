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
	defines   map[string]interface{}
	callbacks map[string]otto.Value
	objects   map[string]otto.Value
}

// Load loads and compiles a plugin given its path with the
// provided definitions.
func Load(path string, defines map[string]interface{}) (*Plugin, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	plugin := &Plugin{
		Name:      strings.Replace(filepath.Base(path), ".js", "", -1),
		Code:      string(raw),
		Path:      path,
		defines:   defines,
		callbacks: make(map[string]otto.Value),
		objects:   make(map[string]otto.Value),
	}

	if err = plugin.compile(); err != nil {
		return nil, err
	}

	for name, val := range defines {
		if err := plugin.vm.Set(name, val); err != nil {
			return nil, err
		}
	}

	return plugin, err
}

// Clone returns a new instance identical to the plugin.
func (p *Plugin) Clone() *Plugin {
	clone, err := Load(p.Path, p.defines)
	if err != nil {
		panic(err) // this should never happen
	}
	return clone
}

// Call executes one of the declared callbacks of the plugin by its name.
func (p *Plugin) Call(name string, args ...interface{}) (interface{}, error) {
	if cb, found := p.callbacks[name]; !found {
		return nil, fmt.Errorf("%s does not name a function", name)
	} else if ret, err := cb.Call(otto.NullValue(), args...); err != nil {
		return nil, err
	} else if !ret.IsUndefined() {
		if exported, err := ret.Export(); err != nil {
			return nil, err
		} else {
			return exported, nil
		}
	}
	return nil, nil
}
