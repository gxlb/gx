package stopwatch

import (
	"fmt"
	"sync"
	"testing"
	"time"
	actime "vipally.gmail.com/basic/c/time"
	//atime "vipally.gmail.com/basic/time"
	//asw "vipally.gmail.com/basic/time/stopwatch"
	"runtime"
	"syscall"
	arand "vipally.gmail.com/basic/math/rand"
)

// // 此类是从 testdll.dll 导出的
// class TESTDLL_API Ctestdll {
// public:
// 	Ctestdll(void);
// 	// TODO: 在此添加您的方法。
// };
// extern TESTDLL_API int ntestdll;
// TESTDLL_API void empty_func(void);
// TESTDLL_API int return_func(int n);

func empty_fun()         {}
func ret_func(n int) int { return n }
func ch_read(ch chan int) {
	var i int
	//t := time.Now()
	//fmt.Println("ch_read begin", t)
	for {
		i = <-ch
		if i < 0 {
			break
		}
	}
	//fmt.Println("ch_read end", time.Now().Sub(t))
}
func ch_write(ch chan int, nCount int) {
	//t := time.Now()
	//fmt.Println("ch_write begin", t)
	for i := 0; i < nCount; i++ {
		ch <- i
	}
	ch <- -1
	//fmt.Println("ch_write end", time.Now().Sub(t))
}

var g_rand uint32 = arand.Rand32()
var g_prime_a32 uint32 = 7368787
var g_prime_c32 uint32 = 2750159

func no_sync_rand() uint32 {
	return g_rand*g_prime_a32 + g_prime_c32
}

func TestTime(t *testing.T) {
	const call_times = 100000

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Some operation time mesure, run times =", call_times)

	const run_cnt = 10
	tt := New()
	for cc := 0; cc < run_cnt; cc++ {
		tt.Restart()

		for i := 0; i < call_times; i++ {
		}
		fmt.Println(tt.ReportWatch("go empty_loop", call_times))

		for i := 0; i < call_times; i++ {
			empty_fun()
		}
		fmt.Println(tt.ReportWatch("go empty_func", call_times))
		for i := 0; i < call_times; i++ {
			ret_func(i)
		}
		fmt.Println(tt.ReportWatch("go with ret", call_times))

		for i := 0; i < call_times; i++ {
			//empty_fun()
			actime.CNoRet()
		}
		fmt.Println(tt.ReportWatch("cgo empty_func", call_times))

		for i := 0; i < call_times; i++ {
			//ret_func(i)
			actime.CWithRet(i)
		}
		fmt.Println(tt.ReportWatch("cgo with ret", call_times))

		for i := 0; i < call_times; i++ {
			actime.RDTSC()
		}
		fmt.Println(tt.ReportWatch("c.RDTSC", call_times))

		for i := 0; i < call_times; i++ {
			actime.Clock()
		}
		fmt.Println(tt.ReportWatch("c.Clock", call_times))

		for i := 0; i < call_times; i++ {
			time.Now()
		}
		fmt.Println(tt.ReportWatch("std_time", call_times))

		////atime.Start(atime.Second / 18)
		//for i := 0; i < call_times; i++ {
		//	Now()
		//}
		////atime.Stop()
		//fmt.Println(tt.ReportWatch("atime", call_times))

		for i := 0; i < call_times; i++ {
			arand.Rand32()
		}
		fmt.Println(tt.ReportWatch("spin lock rand", call_times))

		for i := 0; i < call_times; i++ {
			no_sync_rand()
		}
		fmt.Println(tt.ReportWatch("no_sync_rand", call_times))

		var l sync.Mutex
		for i := 0; i < call_times; i++ {
			l.Lock()
			l.Unlock()
		}
		fmt.Println(tt.ReportWatch("sync.Mutex", call_times))

		var l2 sync.RWMutex
		for i := 0; i < call_times; i++ {
			l2.RLock()
			l2.RUnlock()
		}
		fmt.Println(tt.ReportWatch("sync.RMutex", call_times))

		for i := 0; i < call_times; i++ {
			l2.Lock()
			l2.Unlock()
		}
		fmt.Println(tt.ReportWatch("sync.WMutex", call_times))

		var ch_r, ch_w chan int
		var data, nbuf int
		var chname string

		chname = "ch0"
		ch_r = make(chan int)
		ch_w = make(chan int)
		go ch_read(ch_w)
		for i := 0; i < call_times; i++ {
			ch_w <- i
		}
		ch_w <- -1
		fmt.Println(tt.ReportWatch(chname+" write", call_times))
		go ch_write(ch_r, call_times)
		for {
			data = <-ch_r
			if data < 0 {
				break
			}
		}
		fmt.Println(tt.ReportWatch(chname+" read", call_times))

		buf_lens := []int{1, 2, 3, 5, 10, 100}
		for _, v := range buf_lens {
			nbuf = v
			chname = fmt.Sprintf("ch%d", nbuf)
			ch_r = make(chan int, nbuf)
			ch_w = make(chan int, nbuf)
			go ch_read(ch_w)
			for i := 0; i < call_times; i++ {
				ch_w <- i
			}
			ch_w <- -1
			fmt.Println(tt.ReportWatch(chname+" write", call_times))
			go ch_write(ch_r, call_times)
			for {
				data = <-ch_r
				if data < 0 {
					break
				}
			}
			fmt.Println(tt.ReportWatch(chname+" read", call_times))
			for i := 0; i < call_times; i++ {
				ch_w <- i
				data = <-ch_w
			}
			fmt.Println(tt.ReportWatch(chname+" rw", call_times*2))
		}

		//fmt.Println(syscall.GetLastError())
		for i := 0; i < call_times; i++ {
			syscall.GetLastError()
		}
		fmt.Println(tt.ReportWatch("syscall.GetLastError", call_times))

		//fmt.Println(syscall.ComputerName())
		ttt := uint(call_times / 1000)
		for i := 0; i < int(ttt); i++ {
			syscall.ComputerName()
		}
		fmt.Println(tt.ReportWatch("syscall.ComputerName", ttt))
	}

	w := tt.TellAvgWatch()
	for i, v := range w {
		fmt.Println(v.ReportAvg(i, run_cnt))
	}

	//fmt.Println(tt.Report())

	//if tt.TellStepWatch(0) < tt.TellStepWatch(1) {
	//	t.Errorf("atime.Now() slow than std time")
	//}
}
