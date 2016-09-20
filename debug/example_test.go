package debug_test

import (
	"fmt"
	"vipally.gmail.com/basic/debug"
)

//import "vipally.gmail.com/basic/debug"
func ExampleBt() {
	//example begin
	fmt.Println(debug.Bt())
	fmt.Println(debug.Bts())
	// Output :
	//Func{vipally.gmail.com/basic/debug_test.ExampleBt} FileLine{F:/dev/gocode/src/vipally.gmail.com/basic/debug/example_test.go : 11} pc{42C1DE}
	//#0 Func{vipally.gmail.com/basic/debug_test.ExampleBt} File{F:/dev/gocode/src/vipally.gmail.com/basic/debug/example_test.go : 12} pc{42C24F}
	//#1 Func{testing.runExample} File{C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist465310315/go/src/pkg/testing/example.go : 100} pc{427B1C}
	//#2 Func{testing.RunExamples} File{C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist465310315/go/src/pkg/testing/example.go : 36} pc{427820}
	//#3 Func{testing.Main} File{C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist465310315/go/src/pkg/testing/testing.go : 366} pc{4286B5}
	//#4 Func{main.main} File{C:/DOCUME~1/ally/LOCALS~1/Temp/go-build973468619/vipally.gmail.com/basic/debug/_test/_testmain.go : 48} pc{401180}
	//#5 Func{runtime.main} File{C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist465310315/go/src/pkg/runtime/proc.c : 190} pc{4127EE}
	//#6 Func{runtime.goexit} File{C:/Users/ADMINI~1/AppData/Local/Temp/2/bindist465310315/go/src/pkg/runtime/proc.c : 1223} pc{414480}
}
