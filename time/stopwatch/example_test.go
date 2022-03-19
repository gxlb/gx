package stopwatch_test

import (
	"fmt"
	"time"

	"github.com/gxlb/gx/time/stopwatch"
)

func ExampleSingle() {
	sw := stopwatch.NewSingle(stopwatch.WithUnit(time.Millisecond * 10))
	time.Sleep(time.Second)
	d1 := sw.Duration()
	time.Sleep(time.Second)
	d2 := sw.Restart()
	time.Sleep(time.Second)
	d3 := sw.Duration()
	fmt.Printf("d1=%s\nd2=%s\nd3=%s", d1, d2, d3)

	// Output:
	// d1=1s
	// d2=2s
	// d3=1s
}

func ExampleStopWatch() {
	sw := stopwatch.New(stopwatch.WithStepCacheCap(5),
		stopwatch.WithUnit(stopwatch.Millisecond*10),
		stopwatch.WithName("stat"))
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(100*(i+1)))
		sw.AddStepWatch(fmt.Sprintf("step%d", i+1), (i-2)*2)
	}
	fmt.Println(sw.ReportAll())
	time.Sleep(time.Second)
	fmt.Println(sw.ReportOnce("final", 0))

	// Output:
	// StopWatch(stat):
	// i=1	n=step1	t=100ms	s=100ms
	// i=2	n=step2	t=300ms	s=200ms
	// i=3	n=step3	t=600ms	s=300ms
	// i=4	n=step4	t=1s	s=400ms	u(2)=200ms
	// i=5	n=step5	t=1.5s	s=500ms	u(4)=125ms
	// n=final	t=2.5s	s=1s
}
