package debug

import (
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	s := Bt()
	sesp := "Func{vipally.gmail.com/basic/debug.TestDebug}"
	if !strings.HasPrefix(s, sesp) {
		t.Errorf("debug.Bt() error [%s]", s)
	}

	s = Bts()
	sesp = "#0 Func{vipally.gmail.com/basic/debug.TestDebug}"
	if !strings.HasPrefix(s, sesp) {
		t.Errorf("debug.Bts() error [%s]", s)
	}
}
