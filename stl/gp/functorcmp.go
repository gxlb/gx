//this file is used to import by other gp files
//it cannot use independently, simulation C++ stl functors
package gp

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
// #GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)
// #GOGP_REQUIRE(github.com/vipally/gx/stl/gp/factorcmp)
//
//
///*   //<----This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
//
// This is exactly not a real go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

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

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

//
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

const (
	CMP_Lesser = iota //default
	CMP_Greater
	_CMP_CNT_
)

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePart byte

//create cmp object by name
func CreateCmpGOGPGlobalNamePart(cmpName string) (r CmpGOGPGlobalNamePart) {
	r = CmpGOGPGlobalNamePart(0).CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePart) F(left, right GOGPValueType) (ok bool) {
	switch me {
	case CMP_Lesser:
		ok = me.less(left, right)
	case CMP_Greater:
		ok = me.great(left, right)
	}
	return
}

func (me CmpGOGPGlobalNamePart) Lesser() CmpGOGPGlobalNamePart  { return CMP_Lesser }
func (me CmpGOGPGlobalNamePart) Greater() CmpGOGPGlobalNamePart { return CMP_Greater }

func (me CmpGOGPGlobalNamePart) String() (s string) {
	switch me {
	case CMP_Lesser:
		s = "Lesser"
	case CMP_Greater:
		s = "Greater"
	default:
		s = "error"
	}
	return
}

func (me CmpGOGPGlobalNamePart) CreateByName(cmpName string) (r CmpGOGPGlobalNamePart) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = CMP_Lesser
	case "Greater":
		r = CMP_Greater
	default: //unsupport name
		panic(cmpName)
	}
	return
}

func (me CmpGOGPGlobalNamePart) less(left, right GOGPValueType) (ok bool) {
	//#GOGP_IFDEF GOGP_HasCmpFunc
	ok = left.Less(right)
	//#GOGP_ELSE
	ok = left < right
	//#GOGP_ENDIF
	return
}
func (me CmpGOGPGlobalNamePart) great(left, right GOGPValueType) (ok bool) {
	//#GOGP_IFDEF GOGP_HasCmpFunc
	ok = right.Less(left)
	//#GOGP_ELSE
	ok = right < left
	//#GOGP_ENDIF
	return
}

////////////////////////////////////////////////////////////////////////////////
//type ComparerGOGPGlobalNamePart interface {
//	F(left, right GOGPValueType) bool
//}

////create cmp object by name
//func CreateComparerGOGPGlobalNamePart(cmpName string) (r ComparerGOGPGlobalNamePart) {
//	switch cmpName {
//	case "": //default Lesser
//		fallthrough
//	case "Lesser":
//		r = LesserGOGPGlobalNamePart{}
//	case "Greater":
//		r = GreaterGOGPGlobalNamePart{}
//	default: //unsupport name
//		panic(cmpName)
//	}
//	return
//}

////Lesser
//type LesserGOGPGlobalNamePart struct{}

//func (this LesserGOGPGlobalNamePart) F(left, right GOGPValueType) (ok bool) {
//	//#GOGP_IFDEF GOGP_HasCmpFunc
//	ok = left.Less(right)
//	//#GOGP_ELSE
//	ok = left < right
//	//#GOGP_ENDIF
//	return
//}

////Greater
//type GreaterGOGPGlobalNamePart struct{}

//func (this GreaterGOGPGlobalNamePart) F(left, right GOGPValueType) (ok bool) {
//	//#GOGP_IFDEF GOGP_HasCmpFunc
//	ok = right.Less(left)
//	//#GOGP_ELSE
//	ok = left > right
//	//#GOGP_ENDIF
//	return
//}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
