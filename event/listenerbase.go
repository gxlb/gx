package event

type ListenerBase struct {
	disable      bool
	handler      Handler
	listenerId   ListenerId
	listenerName string
}
