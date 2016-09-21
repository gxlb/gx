package errors

type withidErrors errId

//new error only with id WithIdErr
func New(_id errId) WithIdErr {
	return withidErrors(_id)
}

func (me withidErrors) Error() string {
	return errId(me).String()
}

func (me withidErrors) Id() errId {
	return errId(me)
}
