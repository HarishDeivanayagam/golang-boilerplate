package di

import "sync"

type container map[string]interface{}

// singleton
var dependency container

var lock = &sync.Mutex{}

// UseIOC instanciated DI map
func UseIOC() {
	lock.Lock()
	defer lock.Unlock()
	if dependency == nil {
		dependency = make(map[string]interface{})
	}
}

// AddDependency to the DI Container
func AddDependency(ref string, impl interface{}) {
	if dependency != nil {
		dependency[ref] = impl
	}
}

// Inject injects a dependency
func Inject(req string) interface{} {
	if dependency != nil {
		return dependency[req]
	}
	return nil

}
