package event

import (
	"github.com/vipally/gx/event/eventbase"
)

//define an event
func DefineEvent(id int, name string) eventbase.EventId {
	return 0
}

//register an event creater
func RegEvent(creater eventbase.Creater) error {
	return nil
}
