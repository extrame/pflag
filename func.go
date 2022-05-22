package pflag

import "fmt"

// -- string Value
type funcValue Fn

type Fn func(value string) error

func newFuncValue(val Fn) funcValue {
	return funcValue(val)
}

func (s funcValue) Set(val string) error {
	return s(val)
}

func (s funcValue) String() string { return fmt.Sprintf("%v", Fn(s)) }

// Func defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) Func(name string, value Fn, usage string) {
	f.VarP(newFuncValue(value), name, "", usage)
}

// Like String, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) FuncP(name, shorthand string, value Fn, usage string) {
	f.VarP(newFuncValue(value), name, shorthand, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func Func(name string, value Fn, usage string) {
	CommandLine.FuncP(name, "", value, usage)
}

// Like String, but accepts a shorthand letter that can be used after a single dash.
func FuncP(name, shorthand string, value Fn, usage string) {
	CommandLine.FuncP(name, shorthand, value, usage)
}
