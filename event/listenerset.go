package event

type ListernerSet struct {
	lastId ListenerId //ID generator
	list   []Listener
}
