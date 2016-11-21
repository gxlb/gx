//this file define a template type for sort

package gp

//#GOGP_FILE_BEGIN

import "sort"

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/functorcmp)

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
