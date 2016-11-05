///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sat Nov 05 2016 20:36:39]
// Generate from:
//   [github.com/vipally/gx/fs/tree.gp_filehash38b3.go]
//   [github.com/vipally/gx/fs/tree.gpg] [tree_filehash]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
//
///////////////////////////////////////////////////////////////////

//this file is used to import by other gp files
//it cannot use independently, simulation C++ stl functors
package fs

//#GOGP_ONCE
const (
	CMPLesser = iota //default
	CMPGreater
) //
//#GOGP_ONCE_END

//cmp object, zero is Lesser
type CmpFileHash byte

const (
	CmpFileHashLesser  CmpFileHash = CMPLesser
	CmpFileHashGreater CmpFileHash = CMPGreater
)

//create cmp object by name
func CreateCmpFileHash(cmpName string) (r CmpFileHash) {
	r = CmpFileHashLesser.CreateByName(cmpName)
	return
}

//uniformed global function
func (me CmpFileHash) F(left, right *FileHashTreeNode) (ok bool) {
	switch me {
	case CMPLesser:
		ok = me.less(left, right)
	case CMPGreater:
		ok = me.great(left, right)
	}
	return
}

//Lesser object
func (me CmpFileHash) Lesser() CmpFileHash { return CMPLesser }

//Greater object
func (me CmpFileHash) Greater() CmpFileHash { return CMPGreater }

//show as string
func (me CmpFileHash) String() (s string) {
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
func (me CmpFileHash) CreateByBool(bigFirst bool) (r CmpFileHash) {
	if bigFirst {
		r = CMPGreater
	} else {
		r = CMPLesser
	}
	return
}

//create cmp object by name
func (me CmpFileHash) CreateByName(cmpName string) (r CmpFileHash) {
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
func (me CmpFileHash) less(left, right *FileHashTreeNode) (ok bool) {

	ok = left.Less(right)

	return
}

//Greater operation
func (me CmpFileHash) great(left, right *FileHashTreeNode) (ok bool) {

	ok = right.Less(left)

	return
}