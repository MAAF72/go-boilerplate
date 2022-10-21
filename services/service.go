package services

import (
	"context"

	"github.com/MAAF72/go-boilerplate/adapters"
	"github.com/MAAF72/go-boilerplate/repositories"
)

type ctxKey string

// service const
const (
	DBREPOKEY ctxKey = "db_repositories"
)

type service interface {
	UserService
	ItemService
	TransactionService
	TransactionItemService
}

// Instance service singleton
var Instance service

// Service service struct
type Service struct {
	repositories repositories.Repositories
}

// Init init services
func Init(adapters adapters.Adapters) {
	repositories := repositories.Init(adapters)
	Instance = Service{
		repositories: repositories,
	}
}

// DatabaseRepositories get database repositories
func DatabaseRepositories(ctx context.Context, service Service) repositories.DatabaseRepositoriesImpl {
	ctxDbRepo := ctx.Value(DBREPOKEY)

	if ctxDbRepo != nil {
		if dbRepo, ok := ctxDbRepo.(repositories.DatabaseRepositoriesImpl); ok {
			return dbRepo
		}
	}

	return service.repositories.DatabaseRepository()
}
