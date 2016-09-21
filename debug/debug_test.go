package debug

import (
	"strings"
	"testing"
)

func TestDebug(t *testing.T) {
	s := Bt()
	sesp := " Func{github.com/vipally/gx/debug.TestDebug}"
	if !strings.HasPrefix(s, sesp) {
		t.Errorf("debug.Bt() error \n[%s]\n[%s]", s, sesp)
	}

	s = Bts()
	sesp = "#0 Func{github.com/vipally/gx/debug.TestDebug}"
	if !strings.HasPrefix(s, sesp) {
		t.Errorf("debug.Bts() error [%s]", s)
	}
}
