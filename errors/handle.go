package errors

type ErrHandle byte

const (
	HandleIgnore = iota
	HandleLog
	HandlePanic
)

func (me ErrHandle) Handle() {}
