package errors

import (
	"fmt"
)

type withidstrErrors struct {
	id  errId
	err string
}

//new with id and string WithIdErr
func Newf(_id errId, _fmt string, _para ...interface{}) WithIdErr {
	return &withidstrErrors{id: _id, err: fmt.Sprintf(_fmt, _para...)}
}

func (me *withidstrErrors) Error() string {
	s, err := me.id.Get()
	if err != nil {
		return err.Error()
	}
	s = fmt.Sprintf("error#%d (%s): %s", me.id, s, me.err)
	return s
}

func (me *withidstrErrors) Id() errId {
	return me.id
}
