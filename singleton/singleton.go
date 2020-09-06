package singleton

import "sync"

type manager struct {
	state string
}

var singleton *manager

var once sync.Once

//GetManager is the method to get the state
func GetManager() *manager {
	// once.Do function ensures that the singleton is only instantiated once
	once.Do(func() {
		singleton = &manager{state: "off"}
	})

	return singleton
}

// GetState is a method to get the manager's state
func (sm *manager) GetState() string {
	return sm.state
}

// SetState is a method to set the value of manager's state
func (sm *manager) SetState(s string) {
	sm.state = s
}
