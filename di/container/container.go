package container

import (
	"clean-architecture/di/provider"
	"clean-architecture/infrastructure"
	"clean-architecture/repository"
	"clean-architecture/services"
)

type container struct {
	booksRepository repository.BooksRepository
	booksService    services.BooksService
}

var containerInstance container

func assignMap() map[string](interface{}) {
	injectMap := make(map[string](interface{}))
	injectMap["BooksRepository"] = containerInstance.booksRepository
	injectMap["BooksService"] = containerInstance.booksService
	return injectMap
}

// UseContainer instanciates a container
func UseContainer() {
	db := infrastructure.GetDB()
	containerInstance = container{
		booksRepository: &infrastructure.BooksRepositoryImpl{DB: db},
		booksService:    &services.BooksServiceImpl{},
	}
	provider.InjectMap = assignMap()
}
