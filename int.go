package bindflag

import (
	"flag"
	"fmt"
	"strconv"
)

type uintValue struct {
	p *uint
}

// Uint defines a uint flag that is bound to the provided memory location.
func Uint(p *uint) flag.Value {
	return &uintValue{p}
}

func (v *uintValue) String() string {
	if v.p == nil {
		return "0"
	}
	return fmt.Sprintf("%v", *v.p)
}

func (v *uintValue) Set(s string) (err error) {
	i, err := strconv.ParseUint(s, 0, 64)
	*v.p = uint(i)
	return
}

type intValue struct {
	p *int
}

// Int defines an int flag that is bound to the provided memory location.
func Int(p *int) flag.Value {
	return &intValue{p}
}

func (v *intValue) String() string {
	if v.p == nil {
		return "0"
	}
	return fmt.Sprintf("%v", *v.p)
}

func (v *intValue) Set(s string) (err error) {
	i, err := strconv.ParseInt(s, 0, 64)
	*v.p = int(i)
	return
}
