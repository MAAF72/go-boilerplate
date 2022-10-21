package database

import (
	"github.com/MAAF72/go-boilerplate/models"
)

// Migrate migrate the models
func Migrate(db *Database) (err error) {
	modelList := []interface{}{
		models.User{},
		models.Item{},
		models.Transaction{},
		models.TransactionItem{},
	}

	for _, model := range modelList {
		err = db.AutoMigrate(model)
		if err != nil {
			return
		}
	}

	return
}
