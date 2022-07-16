package test

import "github.com/enosi-rl/gtrace/test/internal"

//go:generate gtrace -v

//gtrace:gen
//gtrace:set shortcut
// NOTE: must compile without unused imports error.
type TraceVariadic struct {
	OnSomethingA func(...Type)
	OnSomethingB func(...string)
	OnSomethingC func(...internal.Type)
	OnSomethingD func(int, string, ...Type) func(...Type)
	OnSomethingE func(bool, float32, ...internal.Type) func(bool, ...internal.Type)
}
