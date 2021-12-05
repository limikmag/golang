package singleton

import "sync"

/****************************************************************************************
Singleton pattern represents a global state and most of the time reduces testability.   *
*****************************************************************************************/

type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
