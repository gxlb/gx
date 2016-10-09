package filelock

//lock using mmap file
type RWMutex struct {
	w Mutex
}

func (this *RWMutex) Lock(path string) {

}

func (this *RWMutex) Unlock() {

}

func (this *RWMutex) RLock(path string) {

}

func (this *RWMutex) RUnlock() {

}
