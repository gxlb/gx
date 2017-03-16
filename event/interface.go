package event

//event handler
type Handler interface {
	OnEvent(para ...interface{}) error
}

//func event handler
type HandlerFunc func(para ...interface{}) error

func (f HandlerFunc) OnEvent(para ...interface{}) error {
	return f(para...)
}

type EventId uint32

func (me EventId) Get() uint32 {
	return uint32(me)
}

type ListenerId uint32

func (me ListenerId) Get() uint32 {
	return uint32(me)
}

//Listener creater
type Creater interface {
	Create(listenerName string, listenerId ListenerId, handler Handler, filters ...interface{}) Listener //
	EventId() EventId                                                                                    //
	EventName() string                                                                                   //
}

//event listener
type Listener interface {
	Creater                            //Listener must create itself
	Filt(para ...interface{}) bool     //if Filt return true,the OnEvent will not be callback
	OnEvent(para ...interface{}) error //event Callback func
	ListenerId() ListenerId            //
	ListenerName() string              //
	IsEnable() bool                    //
	Enable(state bool) bool            //
}
