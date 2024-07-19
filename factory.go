package writer

import (
	"fmt"
	"io"
	"sort"

	"github.com/gopi-frame/collection/kv"
	"github.com/gopi-frame/contract/writer"
	"github.com/gopi-frame/exception"
)

var drivers = kv.NewMap[string, writer.Driver]()

// Register register driver
func Register(driverName string, driver writer.Driver) {
	drivers.Lock()
	defer drivers.Unlock()
	if drivers.ContainsKey(driverName) {
		panic(exception.NewArgumentException("driverName", driverName, fmt.Sprintf("duplicate driver \"%s\"", driverName)))
	}
	drivers.Set(driverName, driver)
}

// Drivers list registered drivers
func Drivers() []string {
	drivers.RLock()
	defer drivers.RUnlock()
	list := drivers.Keys()
	sort.Strings(list)
	return list
}

// Open opens writer
func Open(driverName string, options map[string]any) (io.WriteCloser, error) {
	drivers.RLock()
	driver, ok := drivers.Get(driverName)
	drivers.RUnlock()
	if !ok {
		return nil, exception.NewArgumentException("driverName", driverName, fmt.Sprintf("unknown driver \"%s\"", driverName))
	}
	return driver.Open(options)
}
