package pflag

import (
	"fmt"
	"strconv"
)

// -- int16 Value
type int16Value int16

func newInt16Value(val int16, p *int16) *int16Value {
	*p = val
	return (*int16Value)(p)
}

func (i *int16Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 16)
	*i = int16Value(v)
	return err
}

func (i *int16Value) SetInt(v int64) {
	*i = int16Value(v)
}

func (i *int16Value) String() string { return fmt.Sprintf("%v", *i) }

// Int16Var defines an int16 flag with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag.
func (f *FlagSet) Int16Var(p *int16, name string, value int16, usage string) {
	f.VarP(newInt16Value(value, p), name, "", usage)
}

// Like Int16Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
	f.VarP(newInt16Value(value, p), name, shorthand, usage)
}

// Int16Var defines an int16 flag with specified name, default value, and usage string.
// The argument p points to an int16 variable in which to store the value of the flag.
func Int16Var(p *int16, name string, value int16, usage string) {
	CommandLine.VarP(newInt16Value(value, p), name, "", usage)
}

// Like Int16Var, but accepts a shorthand letter that can be used after a single dash.
func Int16VarP(p *int16, name, shorthand string, value int16, usage string) {
	CommandLine.VarP(newInt16Value(value, p), name, shorthand, usage)
}

// Int16 defines an int16 flag with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag.
func (f *FlagSet) Int16(name string, value int16, usage string) *int16 {
	p := new(int16)
	f.Int16VarP(p, name, "", value, usage)
	return p
}

// Like Int16, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Int16P(name, shorthand string, value int16, usage string) *int16 {
	p := new(int16)
	f.Int16VarP(p, name, shorthand, value, usage)
	return p
}

// Int16 defines an int16 flag with specified name, default value, and usage string.
// The return value is the address of an int16 variable that stores the value of the flag.
func Int16(name string, value int16, usage string) *int16 {
	return CommandLine.Int16P(name, "", value, usage)
}

// Like Int16, but accepts a shorthand letter that can be used after a single dash.
func Int16P(name, shorthand string, value int16, usage string) *int16 {
	return CommandLine.Int16P(name, shorthand, value, usage)
}
