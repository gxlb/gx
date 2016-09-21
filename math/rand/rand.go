//Package rand implements some useful rand object
package rand

import (
	"time"

	xtime "github.com/vipally/gx/time"
)

//generate a seed for rand
func RandSeed(init uint64) uint64 {
	r := xtime.RDTSC() + uint64(time.Now().Unix()) + init
	r = r*4294967291 + 8615693
	return r
}
