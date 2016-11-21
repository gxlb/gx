//#GOGP_IFDEF GOGP_DO_NOT_HAS_THIS_DEFINE__
//this is a fake types for other gp file //#GOGP_ENDIF

//#GOGP_IGNORE_BEGIN
package gp //#GOGP_IGNORE_END

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

//these defines are used to make sure this fake go file can be compiled correctlly
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
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
