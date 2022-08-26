package design_pattern

import "sync"

type Singleton struct{}

var single *Singleton
var lock sync.Mutex

func (s *Singleton) Init() {
	single = &Singleton{}
}

func (s *Singleton) GetInstance() *Singleton {
	lock.Lock()
	defer lock.Unlock()

	if single == nil {
		return single
	} else {
		return &Singleton{}
	}
}
