package repositories

import (
	"github.com/MAAF72/go-boilerplate/adapters"
)

// DatabaseRepositoriesImpl Database repositories implementations
type DatabaseRepositoriesImpl interface {
	UserRepository
	ItemRepository
	TransactionRepository
	TransactionItemRepository
	Transaction(operation func(DatabaseRepositoriesImpl) error) error
}

// Repositories Repositories interface
type Repositories interface {
	DatabaseRepository() DatabaseRepositoriesImpl
	GruleRepository() GruleRepositoryImpl
}

type repositories struct {
	dbRepo    DatabaseRepository
	gruleRepo GruleRepository
}

func (repo repositories) DatabaseRepository() DatabaseRepositoriesImpl {
	return repo.dbRepo
}

func (repo repositories) GruleRepository() GruleRepositoryImpl {
	return repo.gruleRepo
}

// Init init repositories
func Init(adapters adapters.Adapters) Repositories {
	return repositories{
		dbRepo:    DatabaseRepository{db: adapters.Database},
		gruleRepo: GruleRepository{grule: adapters.Grule},
	}
}
