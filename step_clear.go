package main

import (
	"context"
)

type stepClear struct {
	*plugin
}

func newClearStep(plugin *plugin) *stepClear {
	return &stepClear{
		plugin: plugin,
	}
}

func (s *stepClear) Runnable() bool {
	return s.Clear
}

func (s *stepClear) Run(_ context.Context) (err error) {
	return
}
