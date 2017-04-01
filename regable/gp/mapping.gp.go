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

import (
	"fmt"
)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPKeyType int                             //
func (this GOGPKeyType) Less(o GOGPKeyType) bool { return this < o }
func (this GOGPKeyType) Show() string            { return "" } //

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END ///require end from(github.com/vipally/gogp/lib/fakedef)

//key/value mapping
type GOGPGlobalNamePrefixMapping struct {
	normal  map[GOGPKeyType]GOGPValueType
	reverse map[GOGPValueType]GOGPKeyType
}

type GOGPGlobalNamePrefixMappingIdType GOGPKeyType

func NewGOGPGlobalNamePrefixMapping() *GOGPGlobalNamePrefixMapping {
	p := &GOGPGlobalNamePrefixMapping{}
	p.Init()
	return p
}

func (this *GOGPGlobalNamePrefixMapping) Init() {
	this.normal = make(map[GOGPKeyType]GOGPValueType)
	this.reverse = make(map[GOGPValueType]GOGPKeyType)
}

//make mapping
func (this *GOGPGlobalNamePrefixMapping) Insert(k GOGPKeyType, v GOGPValueType) (id GOGPGlobalNamePrefixMappingIdType, err error) {
	if _, ok := this.Find(k); ok {
		err = fmt.Errorf("dumplicate key : %#v", k)
		return
	}
	if _, ok := this.ReverseFind(v); ok {
		err = fmt.Errorf("dumplicate value : %#v", v)
		return
	}
	this.normal[k] = v
	this.reverse[v] = k
	return GOGPGlobalNamePrefixMappingIdType(k), nil
}

//remove by key
func (this *GOGPGlobalNamePrefixMapping) RemoveByKey(k GOGPKeyType) (ok bool) {
	if v, find := this.Find(k); find {
		ok = find
		delete(this.normal, k)
		delete(this.reverse, v)
	}
	return
}

//remove by value
func (this *GOGPGlobalNamePrefixMapping) RemoveByValue(v GOGPValueType) (ok bool) {
	if k, find := this.ReverseFind(v); find {
		ok = find
		delete(this.normal, k)
		delete(this.reverse, v)
	}
	return
}

//find by key
func (this *GOGPGlobalNamePrefixMapping) Find(k GOGPKeyType) (v GOGPValueType, ok bool) {
	v, ok = this.normal[k]
	return
}

//find by value
func (this *GOGPGlobalNamePrefixMapping) ReverseFind(v GOGPValueType) (k GOGPKeyType, ok bool) {
	k, ok = this.reverse[v]
	return
}

//clear all mapping
func (this *GOGPGlobalNamePrefixMapping) Clear() {
	this.Init()
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
