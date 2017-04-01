//this file define a template type for sort

package gp

//#GOGP_FILE_BEGIN
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and cmomile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

import "sort"

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END ///require end from(github.com/vipally/gogp/lib/fakedef)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/vipally/gx/stl/gp/functorcmp)
//this file is used to //import by other gp files
//it cannot use independently, simulation C++ stl functors

//package gp

const (
	CMPLesser = iota //default
	CMPGreater
) //

//cmp object, zero is Lesser
type CmpGOGPGlobalNamePrefix byte

const (
	CmpGOGPGlobalNamePrefixLesser  CmpGOGPGlobalNamePrefix = CMPLesser
	CmpGOGPGlobalNamePrefixGreater CmpGOGPGlobalNamePrefix = CMPGreater
)

//create cmp object by name
func CreateCmpGOGPGlobalNamePrefix(cmpName string) (r CmpGOGPGlobalNamePrefix) {
	r = CmpGOGPGlobalNamePrefixLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpGOGPGlobalNamePrefix) F(left, right GOGPValueType) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpGOGPGlobalNamePrefix) Lesser() CmpGOGPGlobalNamePrefix { return CMPLesser }

//Greater object
func (me CmpGOGPGlobalNamePrefix) Greater() CmpGOGPGlobalNamePrefix { return CMPGreater }

//show as string
func (me CmpGOGPGlobalNamePrefix) String() (s string) {
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

//create by bool
func (me CmpGOGPGlobalNamePrefix) CreateByBool(bigFirst bool) (r CmpGOGPGlobalNamePrefix) {
	if bigFirst {
		r = CMPGreater
	} else {
		r = CMPLesser
	}
	return
}

//create cmp object by name
func (me CmpGOGPGlobalNamePrefix) CreateByName(cmpName string) (r CmpGOGPGlobalNamePrefix) {
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
func (me CmpGOGPGlobalNamePrefix) less(left, right GOGPValueType) (ok bool) {

	ok = left < right

	return
}

//Greater operation
func (me CmpGOGPGlobalNamePrefix) great(left, right GOGPValueType) (ok bool) {

	ok = right < left

	return
}

//#GOGP_IGNORE_END ///require end from(github.com/vipally/gx/stl/gp/functorcmp)

////////////////////////////////////////////////////////////////////////////////

var gGOGPGlobalNamePrefixSortSliceGbl struct {
	cmp CmpGOGPGlobalNamePrefix
}

func init() {
	gGOGPGlobalNamePrefixSortSliceGbl.cmp = gGOGPGlobalNamePrefixSortSliceGbl.cmp.CreateByName("#GOGP_GPGCFG(GOGP_DefaultCmpType)")
}

//new sort object
func NewGOGPGlobalNamePrefixSortSlice(capacity int) *GOGPGlobalNamePrefixSortSlice {
	p := &GOGPGlobalNamePrefixSortSlice{}
	p.Init(capacity)
	return p
}

//sort slice
type GOGPGlobalNamePrefixSortSlice struct {
	d []GOGPValueType
}

//init
func (this *GOGPGlobalNamePrefixSortSlice) Init(capacity int) {
	this.d = make([]GOGPValueType, 0, capacity)
}

//sort
func (this *GOGPGlobalNamePrefixSortSlice) Sort() {
	sort.Sort(this)
}

//data buffer
func (this *GOGPGlobalNamePrefixSortSlice) Buffer() []GOGPValueType {
	return this.d
}

//push
func (this *GOGPGlobalNamePrefixSortSlice) Push(v GOGPValueType) int {
	this.d = append(this.d, v)
	return this.Len()
}

//insert
func (this *GOGPGlobalNamePrefixSortSlice) Insert(v GOGPValueType, idx int) int {
	if idx >= 0 && idx < this.Len() {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], v)
		this.d = append(this.d, right...)
	} else {
		this.d = append(this.d, v)
	}
	return this.Len()
}

//remove
func (this *GOGPGlobalNamePrefixSortSlice) Remove(idx int) (r GOGPValueType, ok bool) {
	if r, ok = this.Get(idx); ok {
		right := this.d[idx+1:]
		this.d = append(this.d[:idx], right...)
	}
	return
}

//pop
func (this *GOGPGlobalNamePrefixSortSlice) Pop() (r GOGPValueType, ok bool) {
	if ok = len(this.d) > 0; ok {
		r = (this.d)[len(this.d)-1]
	}
	this.d = (this.d)[:len(this.d)-1]
	return
}

//get
func (this *GOGPGlobalNamePrefixSortSlice) Get(idx int) (r GOGPValueType, ok bool) {
	if ok = idx >= 0 && idx < this.Len(); ok {
		r = this.d[idx]
	}
	return
}

//must get
func (this *GOGPGlobalNamePrefixSortSlice) MustGet(idx int) (r GOGPValueType) {
	ok := false
	if r, ok = this.Get(idx); !ok {
		panic(idx)
	}
	return
}

//len
func (this *GOGPGlobalNamePrefixSortSlice) Len() int {
	return len(this.d)
}

//sort by Hash decend,the larger one first
func (this *GOGPGlobalNamePrefixSortSlice) Less(i, j int) (ok bool) {
	l, r := (this.d)[i], (this.d)[j]
	return gGOGPGlobalNamePrefixSortSliceGbl.cmp.F(l, r)
}

//swap
func (this *GOGPGlobalNamePrefixSortSlice) Swap(i, j int) {
	(this.d)[i], (this.d)[j] = (this.d)[j], (this.d)[i]
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
