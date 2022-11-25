package test

import "github.com/enosi/gtrace/test/internal"

//go:generate gtrace -v

// NOTE: must compile without unused imports error.
//
//gtrace:gen
//gtrace:set shortcut
type TraceArray struct {
	OnSomethingA func(Type)
	OnSomethingB func([]internal.Type)
}
