// Code generated by gtrace. DO NOT EDIT.

// +build !gtrace

package main

// Compose returns a new Trace which has functional fields composed
// both from t and x.
func (t Trace) Compose(x Trace) (ret Trace) {
	switch {
	case t.OnInput == nil:
		ret.OnInput = x.OnInput
	case x.OnInput == nil:
		ret.OnInput = t.OnInput
	default:
		h1 := t.OnInput
		h2 := x.OnInput
		ret.OnInput = func(in0 string) func() {
			r1 := h1(in0)
			r2 := h2(in0)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func() {
					r1()
					r2()
				}
			}
		}
	}
	return ret
}
func gtraceNoop() {
}
func (Trace) onInput(string) func() {
	return gtraceNoop
}
func gtraceNoop1() {
}
func traceOnInput(Trace, string) func() {
	return gtraceNoop1
}