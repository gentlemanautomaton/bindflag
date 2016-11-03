package bindflag

import (
	"flag"
	"time"
)

type durationValue struct {
	p *time.Duration
}

// Duration defines a duration flag that is bound to the provided memory
// location.
func Duration(duration *time.Duration) flag.Value {
	return &durationValue{duration}
}

func (v *durationValue) String() string {
	if v.p == nil {
		return time.Duration(0).String()
	}
	return v.p.String()
}

func (v *durationValue) Set(value string) (err error) {
	*v.p, err = time.ParseDuration(value)
	return
}

/*
type seconds struct {
  d *time.Duration
}

// Seconds returns a flags.Value that interprets interprets a string value as
// a number of seconds and saves that value in the provided duration.
func Seconds(duration *time.Duration) flags.Value {
  return &seconds{duration}
}

func (s *seconds) String() string {
	return s.d.
}

func (u *uintOrInf) Set(value string) (err error) {
  dur, d1 := time.ParseDuration()
	u.Present = true

	switch strings.ToLower(value) {
	case "infinite", "inf", "i":
		u.Inf = true
	default:
		v, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return fmt.Errorf("\"%s\" is neither an acceptable number nor \"infinite\"", value)
		}
		u.Value = uint(v)
	}

	return
}
*/
