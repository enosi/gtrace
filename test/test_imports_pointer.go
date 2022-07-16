package test

import "github.com/enosi-rl/gtrace/test/internal"

//go:generate gtrace -v

//gtrace:gen
//gtrace:set shortcut
// NOTE: must compile without unused imports error.
type TracePointer struct {
	OnSomethingA func(Type)
	OnSomethingB func(*internal.Type)
}
