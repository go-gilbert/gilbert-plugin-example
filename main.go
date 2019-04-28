package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/x1unix/gilbert/log"
	"github.com/x1unix/gilbert/manifest"
	"github.com/x1unix/gilbert/plugins"
	"github.com/x1unix/gilbert/scope"
)

func NewPlugin(scope *scope.Scope, p manifest.RawParams, log log.Logger) (plugins.Plugin, error) {
	params := Params{}
	if err := mapstructure.Decode(params, &p); err != nil {
		return nil, fmt.Errorf("failed to read configuration: %s", err)
	}

	return &Plugin{
		scope:  scope,
		params: params,
	}, nil
}

func main() {}
