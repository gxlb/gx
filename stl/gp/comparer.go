//import from github.com/vipally/gx/stl/gp/comparer.go begin
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

//GOGP_IGNORE_BEGIN
//this file is used to import by other gp files
//it cannot use independently
//simulation C++ stl functors
package gp //
//GOGP_IGNORE_END

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
//
//
/*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is exactly not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int

func (this GOGPValueType) Less(o GOGPValueType) bool {
	return this < o
}

//
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

type Lesser int

func (this *Lesser) F(left, right GOGPValueType) (ok bool) {
	//#if GOGP_HasLess
	ok = left.Less(right)
	//#else
	ok = left < right
	//#endif
	return
}

//GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
