//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Oct 30 2016 21:17:28]
// Generate from:
//   [github.com/vipally/gx/stl/gp/funccmp.go]
//   [github.com/vipally/gx/stl/gp/gp.gpg] [GOGP_REVERSE_funccmp]
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
//#GOGP_IGNORE_END

//this file is used to import by other gp files
//it cannot use independently, simulation C++ stl functors
<PACKAGE>

//#GOGP_REQUIRE(github.com/vipally/gx/stl/gp/fakedef,_)

type CmpFunc<GLOBAL_NAME_PART> func(left, right <VALUE_TYPE>) bool

//create cmp object by name
func GetCmpFunc<GLOBAL_NAME_PART>(cmpName string) (r CmpFunc<GLOBAL_NAME_PART>) {
	switch cmpName {
	case "": //default Lesser
		fallthrough
	case "Lesser":
		r = Less<GLOBAL_NAME_PART>
	case "Greater":
		r = Great<GLOBAL_NAME_PART>
	default: //unsupport name
		panic(cmpName)
	}
	return
}

//Lesser
func Less<GLOBAL_NAME_PART>(left, right <VALUE_TYPE>) (ok bool) {
	//#GOGP_IFDEF GOGP_HasCmpFunc
	ok = left.Less(right)
	//#GOGP_ELSE
	ok = left < right
	//#GOGP_ENDIF
	return
}

//Greater
func Great<GLOBAL_NAME_PART>(left, right <VALUE_TYPE>) (ok bool) {
	//#GOGP_IFDEF GOGP_HasCmpFunc
	ok = right.Less(left)
	//#GOGP_ELSE
	ok = left > right
	//#GOGP_ENDIF
	return
}

