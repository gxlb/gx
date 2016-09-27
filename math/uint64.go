package math

import (
	"fmt"
)

type Uint64 uint64

func (me Uint64) CommaString() string {
	r := Uint64(1000)
	l := me
	s := ""
	for l >= r {
		d := l % r
		l /= r
		s = fmt.Sprintf(",%03d%s", d, s)
	}
	s = fmt.Sprintf("%d%s", l, s)
	return s
}
