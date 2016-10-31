package gp

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
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
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//import here...
//

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/fakedef)
//these defines is used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/fakedef)



//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gx/stl/gp/functorcmp)
//this file is used to import by other gp files
//it cannot use independently, simulation C++ stl functors
//package gp

//#GOGP_ONCE
const (
	CMPLesser = iota //default
	CMPGreater
) //
//#GOGP_ONCE_END

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePart byte

const (
	CmpGOGPGlobalNamePartLesser  CmpGOGPGlobalNamePart = CMPLesser
	CmpGOGPGlobalNamePartGreater CmpGOGPGlobalNamePart = CMPGreater
)

//create cmp object by name
func CreateCmpGOGPGlobalNamePart(cmpName string) (r CmpGOGPGlobalNamePart) {
	r = CmpGOGPGlobalNamePartLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePart) F(left, right GOGPValueType) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpGOGPGlobalNamePart) Lesser() CmpGOGPGlobalNamePart { return CMPLesser }

//Greater object
func (me CmpGOGPGlobalNamePart) Greater() CmpGOGPGlobalNamePart { return CMPGreater }

//show as string
func (me CmpGOGPGlobalNamePart) String() (s string) {
	switch me {
	case CMPLesser:
		s = "Lesser"
	case CMPGreater:
		s = "Greater"
	default:
		s = "error cmp value"
	}
	return
}

//create cmp object by name
func (me CmpGOGPGlobalNamePart) CreateByName(cmpName string) (r CmpGOGPGlobalNamePart) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = CMPLesser
	case "Greater":
		r = CMPGreater
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//lesser operation
func (me CmpGOGPGlobalNamePart) less(left, right GOGPValueType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePart) great(left, right GOGPValueType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END //required from(github.com/vipally/gx/stl/gp/functorcmp)

////#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
//
//
//
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

type GOGPGlobalNamePrefixBSTree struct {
	root *GOGPGlobalNamePrefixTreeNode
}

type GOGPGlobalNamePrefixBSTreeNode struct{}

//
//add file body here..........
//
//

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
