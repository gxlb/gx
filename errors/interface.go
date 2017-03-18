package errors

//id of error,an ErrorId is a predefined error type
type ErrorId uint32

type Error interface {
	error                          //std error interface
	New(para ...interface{}) Error //create another instance, call Bts()?
	Handle()                       //handle this error
	Id() ErrorId                   //Id of this error
}

type ErrorObj struct {
	Id   ErrorId
	Para []interface{}
}

//regist an error
func Reg(moduleId uint16, name, content string, handle Handler) ErrorId {
	return 0
}

type ErrorInfo struct {
	content string
	name    string
	id      ErrorId
	ref     int
	Objs    []ErrorObj
}

type Handler byte

const (
	HandleIgnore Handler = 1 << iota
	HandleLog
	HandlePrint
	HandlePanic
	HandleBts
)

func (me Handler) Handle() {}
