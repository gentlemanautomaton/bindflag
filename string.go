package bindflag

import "flag"

type str struct {
	p *string
}

// String defines a string flag that is bound to the provided memory location.
func String(p *string) flag.Value {
	return &str{p}
}

func (v *str) String() string {
	if v.p == nil {
		return ""
	}
	return *v.p
}

func (v *str) Set(value string) (err error) {
	*v.p = value
	return
}
