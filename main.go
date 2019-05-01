package main

import (
	"github.com/go-gilbert/gilbert-sdk"
)

// NewHelloAction is action constructor
func NewHelloAction(scope sdk.ScopeAccessor, params sdk.ActionParams) (sdk.ActionHandler, error) {
	p := Params{}
	if err := params.Unmarshal(&p); err != nil {
		return nil, err
	}

	return &HelloAction{
		scope:  scope,
		params: p,
	}, nil
}

// GetPluginName returns plugin name
func GetPluginName() string {
	return "example-plugin"
}

// GetPluginActions returns available plugin actions
func GetPluginActions() sdk.Actions {
	return sdk.Actions{
		"hello-world": NewHelloAction,
	}
}
