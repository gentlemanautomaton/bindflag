package bindflag

import (
	"flag"
	"fmt"
	"strconv"
)

type boolValue struct {
	p *bool
}

// Bool defines a bool flag that is bound to the provided memory location.
func Bool(p *bool) flag.Value {
	return &boolValue{p}
}

func (v *boolValue) String() string {
	if v.p == nil {
		return "false"
	}
	return fmt.Sprintf("%v", *v.p)
}

func (v *boolValue) Set(s string) (err error) {
	c, err := strconv.ParseBool(s)
	*v.p = bool(c)
	return
}

func (v *boolValue) IsBoolFlag() bool { return true }
