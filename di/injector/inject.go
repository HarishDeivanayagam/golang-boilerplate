package injector

import (
	"clean-architecture/di/provider"
)

// Inject injects the requested dependency
func Inject(req string) interface{} {
	return provider.InjectMap[req]
}
