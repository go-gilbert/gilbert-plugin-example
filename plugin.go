package main

import (
	"github.com/x1unix/gilbert/plugins"
	"github.com/x1unix/gilbert/runner/job"
	"github.com/x1unix/gilbert/scope"
)

type Params struct {
	Message string
}

type Plugin struct {
	scope  *scope.Scope
	params Params
}

func (p *Plugin) Call(ctx *job.RunContext, r plugins.JobRunner) (err error) {
	ctx.Logger.Info("Hello World")
	return nil
}

func (p *Plugin) Cancel(ctx *job.RunContext) error {
	ctx.Logger.Info("Cancel callback call")
	return nil
}
