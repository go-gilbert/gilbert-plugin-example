package main

import (
	"github.com/go-gilbert/gilbert-sdk"
)

func NewPlugin(scope sdk.ScopeAccessor, params sdk.PluginParams, _ sdk.Logger) (sdk.Plugin, error) {
	p := Params{}
	if err := params.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &Plugin{
		scope:  scope,
		params: p,
	}, nil
}

func GetPluginName() string {
	return "@go-gilbert/example"
}
