package di

var dependency map[string]interface{}

// UseIOC instanciated DI map
func UseIOC() {
	dependency = make(map[string]interface{})
}

// AddDependency to the DI Container
func AddDependency(ref string, impl interface{}) {
	dependency[ref] = impl
}

// Inject injects a dependency
func Inject(req string) interface{} {
	return dependency[req]
}
