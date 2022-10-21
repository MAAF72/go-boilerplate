package repositories

import (
	"github.com/MAAF72/go-boilerplate/adapters/database"
)

// DatabaseRepository database repository
type DatabaseRepository struct {
	db *database.Database
}

// Transaction start transaction
func (repo DatabaseRepository) Transaction(operation func(DatabaseRepositoriesImpl) error) (err error) {
	err = repo.db.Transaction(func(tx *database.DB) (err error) {
		transactionRepo := DatabaseRepository{
			&database.Database{DB: tx},
		}

		err = operation(transactionRepo)

		return
	})

	return
}
