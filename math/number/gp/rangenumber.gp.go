package gp

//#GOGP_FILE_BEGIN 1

//#GOGP_REQUIRE(github.com/vipally/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN //required from(github.com/vipally/gogp/lib/fakedef)
//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END //required from(github.com/vipally/gogp/lib/fakedef)

type rangeGOGPGlobalNamePrefix struct {
	min, max GOGPValueType
}

func (this *rangeGOGPGlobalNamePrefix) verify(val GOGPValueType) bool {
	return val >= this.min && val <= this.max
}

type GOGPGlobalNamePrefixRangeNumber struct {
	val GOGPValueType
	rng *rangeGOGPGlobalNamePrefix
}

type GOGPGlobalNamePrefix GOGPValueType

func (me GOGPGlobalNamePrefix) Get() GOGPValueType       { return GOGPValueType(me) }
func (this *GOGPGlobalNamePrefix) Set(val GOGPValueType) { *this = GOGPGlobalNamePrefix(val) }
func (me GOGPGlobalNamePrefix) Verify(val GOGPValueType) bool {
	return val >= GOGPGlobalNamePrefixMin && val <= GOGPGlobalNamePrefixMax
}

//#GOGP_FILE_END
