package test

import "github.com/enosi/gtrace/test/internal"

//go:generate gtrace -v

// NOTE: must compile without unused imports error.
//
//gtrace:gen
//gtrace:set shortcut
type TraceVariadic struct {
	OnSomethingA func(...Type)
	OnSomethingB func(...string)
	OnSomethingC func(...internal.Type)
	OnSomethingD func(int, string, ...Type) func(...Type)
	OnSomethingE func(bool, float32, ...internal.Type) func(bool, ...internal.Type)
}
