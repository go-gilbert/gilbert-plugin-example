package main

import (
	"github.com/go-gilbert/gilbert-sdk"
)

type Params struct {
	Message string
}

type Plugin struct {
	scope  sdk.ScopeAccessor
	params Params
}

func (p *Plugin) Call(ctx sdk.JobContextAccessor, r sdk.JobRunner) (err error) {
	ctx.Log().Info(p.params.Message)
	return nil
}

func (p *Plugin) Cancel(ctx sdk.JobContextAccessor) error {
	ctx.Log().Info("Cancel callback call")
	return nil
}
