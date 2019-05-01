package main

import (
	"github.com/go-gilbert/gilbert-sdk"
)

type Params struct {
	Message string
}

type HelloAction struct {
	scope  sdk.ScopeAccessor
	params Params
}

func (p *HelloAction) Call(ctx sdk.JobContextAccessor, r sdk.JobRunner) (err error) {
	ctx.Log().Info(p.params.Message)
	return nil
}

func (p *HelloAction) Cancel(ctx sdk.JobContextAccessor) error {
	ctx.Log().Info("Cancel callback call")
	return nil
}
