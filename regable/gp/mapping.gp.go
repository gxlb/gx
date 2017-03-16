package gp

//#GOGP_FILE_BEGIN

import (
	"fmt"
)

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)

//key/value mapping
type GOGPGlobalNamePrefixMapping struct {
	normal  map[GOGPKeyType]GOGPValueType
	reverse map[GOGPValueType]GOGPKeyType
}

func (this *GOGPGlobalNamePrefixMapping) Init() {
	this.normal = make(map[GOGPKeyType]GOGPValueType)
	this.reverse = make(map[GOGPValueType]GOGPKeyType)
}

//make mapping
func (this *GOGPGlobalNamePrefixMapping) Insert(k GOGPKeyType, v GOGPValueType) error {
	if _, ok := this.Find(k); ok {
		return fmt.Errorf("dumplicate key : %#v", k)
	}
	if _, ok := this.ReverseFind(v); ok {
		return fmt.Errorf("dumplicate value : %#v", v)
	}
	this.normal[k] = v
	this.reverse[v] = k
	return nil
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

//#GOGP_FILE_END
