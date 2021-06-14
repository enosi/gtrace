package test

//go:generate gtrace -v -tag debug

//gtrace:gen
//gtrace:set shortcut
type TraceReturningTrace struct {
	OnReturnedTrace func() ReturnedTrace
}

//gtrace:gen
type ReturnedTrace struct {
	OnSomething func()
}