// export some C time function such as RDTSC() from cgo

package time

/*
#include<time.h>
unsigned long long rdtsc(void)
{
	unsigned long long ullCPUTick = 0;
	__asm__ __volatile__("rdtsc" : "=A" (ullCPUTick) : );
	return ullCPUTick;
}
*/
import "C"

//call rdtsc from C
func RDTSC() uint64 {
	return uint64(C.rdtsc())
}

//call clock() from C
func Clock() uint64 {
	return uint64(C.clock())
}
