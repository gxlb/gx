package rand

import (
	"os"
	"testing"

	"github.com/vipally/gogp"
)

func TestCallGogp(t *testing.T) {
	if dir, err := os.Getwd(); err == nil {
		gogp.Work(dir)
	} else {
		panic(err)
	}
}
