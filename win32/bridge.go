package main

import (
	"C"
	"unsafe"
	"fmt"
	"github.com/x1unix/gilbert-plugin-example"
	"github.com/x1unix/gilbert/log"
	"github.com/x1unix/gilbert/plugins"
	"github.com/x1unix/gilbert/manifest"
	"github.com/x1unix/gilbert/scope"
)

//export NewPlugin
func NewPlugin(scopePtr, paramsPtr, logPtr, resultPtr, errPtr uintptr) {
	err := (*error) (unsafe.Pointer(errPtr))
	defer func() {
		if r := recover(); r != nil {
			*err = fmt.Errorf("panic: %v", r)
		}
	}()
	scp := (*scope.Scope) (unsafe.Pointer(scopePtr))
	params := (*manifest.RawParams) (unsafe.Pointer(paramsPtr))
	logger := (*log.Logger) (unsafe.Pointer(logPtr))
	result := (*plugins.Plugin) (unsafe.Pointer(resultPtr))
	
	plug, plugErr := plugin.NewPlugin(scp, *params, *logger)
	if plugErr != nil {
		*err = plugErr
		return
	}

	*result = plug
}

func main() {}