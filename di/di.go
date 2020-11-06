package di

import "sync"

type container map[string]interface{}

// singleton
var diContainer container

var lock = &sync.Mutex{}

// UseIOC instanciated DI map
func UseIOC() {
	lock.Lock()
	defer lock.Unlock()
	if diContainer == nil {
		diContainer = make(container)
	}
}

// AddDependency to the DI Container
func AddDependency(ref string, impl interface{}) {
	if diContainer != nil {
		diContainer[ref] = impl
	}
}

// Inject injects a dependency
func Inject(req string) interface{} {
	if diContainer != nil {
		return diContainer[req]
	}
	return nil
}
