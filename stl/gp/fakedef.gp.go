//#GOGP_IFDEF GOGP_DO_NOT_HAS_THIS_DEFINE__
//this is a fake types for other gp file //#GOGP_ENDIF

//#GOGP_IGNORE_BEGIN
package gp //#GOGP_IGNORE_END

//#GOGP_FILE_BEGIN

//these defines is used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
//#GOGP_IFDEF KEY_TYPE
type GOGPKeyType int                             //
func (this GOGPKeyType) Less(o GOGPKeyType) bool { return this < o }
func (this GOGPKeyType) Show() string            { return "" } //
//#GOGP_ENDIF //KEY_TYPE

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

//#GOGP_FILE_END
