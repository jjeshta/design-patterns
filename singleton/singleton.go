package singleton

import "sync"

// manager as our singleton object
// Note that we should define the struct with lowercase letters in order to make it private.
// This will disallow instantiating the struct outside the package.
type manager struct {
	state string
}

// Global variable of type pointer to 'manager' that will keep an reference to its singleton instance
var singleton *manager

// declare sync.0nce struct
var once sync.Once

//GetManager provides a global point of access to manager.
// It is in the GetManager where all our expensive code is done.
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
