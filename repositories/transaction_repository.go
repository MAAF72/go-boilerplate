package repositories

import (
	"github.com/MAAF72/go-boilerplate/models"
	"github.com/mitchellh/mapstructure"
)

// TransactionRepository transaction repository interface
type TransactionRepository interface {
	SaveTransaction(request models.Transaction) (id string, err error)
	FindTransactionByID(id string) (transaction models.Transaction, err error)
	UpdateTransactionByID(id string, changeSet models.TransactionChangeSet) (err error)
	DeleteTransactionByID(id string) (err error)
}

// SaveTransaction save transaction
func (repo DatabaseRepository) SaveTransaction(transaction models.Transaction) (id string, err error) {
	err = repo.db.Create(&transaction).Error

	id = transaction.ID

	return
}

// FindTransactionByID find transaction by id
func (repo DatabaseRepository) FindTransactionByID(id string) (transaction models.Transaction, err error) {
	err = repo.db.Take(&transaction, "id = ?", id).Error

	return
}

// UpdateTransactionByID update transaction by id
func (repo DatabaseRepository) UpdateTransactionByID(id string, changeSet models.TransactionChangeSet) (err error) {
	var updates map[string]interface{}
	transaction := models.Transaction{
		Base: models.Base{ID: id},
	}

	mapstructure.Decode(changeSet, &updates)

	err = repo.db.Model(&transaction).Updates(updates).Error

	return
}

// DeleteTransactionByID delete transaction by id
func (repo DatabaseRepository) DeleteTransactionByID(id string) (err error) {
	err = repo.db.Delete(&models.Transaction{}, "id = ?", id).Error

	return
}
