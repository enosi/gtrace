// Code generated by gtrace. DO NOT EDIT.

package test

import (
	"bytes"

	"github.com/enosi-rl/gtrace/test/internal"
)

// Compose returns a new TraceVariadic which has functional fields composed
// both from t and x.
func (t TraceVariadic) Compose(x TraceVariadic) (ret TraceVariadic) {
	switch {
	case t.OnSomethingA == nil:
		ret.OnSomethingA = x.OnSomethingA
	case x.OnSomethingA == nil:
		ret.OnSomethingA = t.OnSomethingA
	default:
		h1 := t.OnSomethingA
		h2 := x.OnSomethingA
		ret.OnSomethingA = func(t ...Type) {
			h1(t...)
			h2(t...)
		}
	}
	switch {
	case t.OnSomethingB == nil:
		ret.OnSomethingB = x.OnSomethingB
	case x.OnSomethingB == nil:
		ret.OnSomethingB = t.OnSomethingB
	default:
		h1 := t.OnSomethingB
		h2 := x.OnSomethingB
		ret.OnSomethingB = func(s ...string) {
			h1(s...)
			h2(s...)
		}
	}
	switch {
	case t.OnSomethingC == nil:
		ret.OnSomethingC = x.OnSomethingC
	case x.OnSomethingC == nil:
		ret.OnSomethingC = t.OnSomethingC
	default:
		h1 := t.OnSomethingC
		h2 := x.OnSomethingC
		ret.OnSomethingC = func(t ...internal.Type) {
			h1(t...)
			h2(t...)
		}
	}
	switch {
	case t.OnSomethingD == nil:
		ret.OnSomethingD = x.OnSomethingD
	case x.OnSomethingD == nil:
		ret.OnSomethingD = t.OnSomethingD
	default:
		h1 := t.OnSomethingD
		h2 := x.OnSomethingD
		ret.OnSomethingD = func(i int, s string, t ...Type) {
			h1(i, s, t...)
			h2(i, s, t...)
		}
	}
	switch {
	case t.OnSomethingE == nil:
		ret.OnSomethingE = x.OnSomethingE
	case x.OnSomethingE == nil:
		ret.OnSomethingE = t.OnSomethingE
	default:
		h1 := t.OnSomethingE
		h2 := x.OnSomethingE
		ret.OnSomethingE = func(b bool, f float32, t ...internal.Type) {
			h1(b, f, t...)
			h2(b, f, t...)
		}
	}
	return ret
}
func (t TraceVariadic) onSomethingA(t1 ...Type) {
	fn := t.OnSomethingA
	if fn == nil {
		return
	}
	fn(t1...)
}
func (t TraceVariadic) onSomethingB(s ...string) {
	fn := t.OnSomethingB
	if fn == nil {
		return
	}
	fn(s...)
}
func (t TraceVariadic) onSomethingC(t1 ...internal.Type) {
	fn := t.OnSomethingC
	if fn == nil {
		return
	}
	fn(t1...)
}
func (t TraceVariadic) onSomethingD(i int, s string, t1 ...Type) {
	fn := t.OnSomethingD
	if fn == nil {
		return
	}
	fn(i, s, t1...)
}
func (t TraceVariadic) onSomethingE(b bool, f float32, t1 ...internal.Type) {
	fn := t.OnSomethingE
	if fn == nil {
		return
	}
	fn(b, f, t1...)
}
func traceVariadicOnSomethingA(t TraceVariadic, e Embeded, s string, integer int, boolean bool, e1 error, r bytes.Reader) {
	var p Type
	p.Embeded = e
	p.String = s
	p.Integer = integer
	p.Boolean = boolean
	p.Error = e1
	p.Reader = r
	t.onSomethingA(p)
}
func traceVariadicOnSomethingB(t TraceVariadic, s string) {
	t.onSomethingB(s)
}
func traceVariadicOnSomethingC(t TraceVariadic) {
	var p internal.Type
	t.onSomethingC(p)
}
func traceVariadicOnSomethingD(t TraceVariadic, i int, s string, e Embeded, s1 string, integer int, boolean bool, e1 error, r bytes.Reader) {
	var p Type
	p.Embeded = e
	p.String = s1
	p.Integer = integer
	p.Boolean = boolean
	p.Error = e1
	p.Reader = r
	t.onSomethingD(i, s, p)
}
func traceVariadicOnSomethingE(t TraceVariadic, b bool, f float32) {
	var p internal.Type
	t.onSomethingE(b, f, p)
}