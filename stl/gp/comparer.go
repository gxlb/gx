//#GOGP_IGNORE_BEGIN
//this file is used to import by other gp files
//it cannot use independently
//simulation C++ stl functors
package gp //
//#GOGP_IGNORE_END

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile_BEGIN
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
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile

//

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPDummyDefine
//
//these defines is used to make sure this dummy go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int

func (this GOGPValueType) Less(o GOGPValueType) bool {
	return this < o
}
func (this GOGPValueType) Great(o GOGPValueType) bool {
	return this > o
}

//
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END////////////////////////////////GOGPDummyDefine

type ComparerGOGPValueType interface {
	F(left, right GOGPValueType) bool
}

type ComparerGOGPValueTypeCreator int

const (
	LESSER_GOGPValueType ComparerGOGPValueTypeCreator = iota
	GREATER_GOGPValueType
)

func (me ComparerGOGPValueTypeCreator) Create() (cmp ComparerGOGPValueType) {
	switch me {
	case LESSER_GOGPValueType:
		cmp = LesserGOGPValueType(0)
	case GREATER_GOGPValueType:
		cmp = GreaterGOGPValueType(0)
	}
	return
}

type LesserGOGPValueType byte

func (this LesserGOGPValueType) F(left, right GOGPValueType) (ok bool) {
	//#GOGP_IFDEF GOGP_HasLess
	ok = left.Less(right)
	//#GOGP_ELSE
	ok = left < right
	//#GOGP_ENDIF
	return
}

type GreaterGOGPValueType byte

func (this GreaterGOGPValueType) F(left, right GOGPValueType) (ok bool) {
	//#GOGP_IFDEF GOGP_HasGreat
	ok = left.Great(right)
	//#GOGP_ELSE
	ok = left > right
	//#GOGP_ENDIF
	return
}

//#GOGP_IGNORE_BEGIN//////////////////////////////GOGPCommentDummyGoFile
//*/
//#GOGP_IGNORE_END////////////////////////////////GOGPCommentDummyGoFile_END
