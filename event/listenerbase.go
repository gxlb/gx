package event

import (
	"github.com/vipally/gx/event/eventbase"
)

type ListenerBase struct {
	disable      bool
	handler      eventbase.Handler
	listenerId   eventbase.ListenerId
	listenerName string
}
