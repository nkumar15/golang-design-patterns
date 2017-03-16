package singleton

import (
	"sync"
)


// My singleton type
type mysingleton struct {
}

var instance *mysingleton
var once sync.Once

// Returns a singleton object of type singleton
func GetInstance() *mysingleton {
	once.Do(func() {
		instance = &mysingleton{}
	})
	return instance
}
