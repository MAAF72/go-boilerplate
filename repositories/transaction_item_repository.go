package repositories

import (
	"github.com/MAAF72/go-boilerplate/models"
	"github.com/mitchellh/mapstructure"
)

// TransactionItemRepository transaction item repository interface
type TransactionItemRepository interface {
	SaveTransactionItem(request models.TransactionItem) (id string, err error)
	FindTransactionItemByID(id string) (transactionItem models.TransactionItem, err error)
	FindAllTransactionItemByTransactionID(transactionID string) (transactionItems []models.TransactionItem, err error)
	UpdateTransactionItemByID(id string, changeSet models.TransactionItemChangeSet) (err error)
	DeleteTransactionItemByID(id string) (err error)
}

// SaveTransactionItem save transaction item
func (repo DatabaseRepository) SaveTransactionItem(transactionItem models.TransactionItem) (id string, err error) {
	err = repo.db.Create(&transactionItem).Error

	id = transactionItem.ID

	return
}

// FindTransactionItemByID find transaction item by id
func (repo DatabaseRepository) FindTransactionItemByID(id string) (transactionItem models.TransactionItem, err error) {
	err = repo.db.Take(&transactionItem, "id = ?", id).Error

	return
}

// FindAllTransactionItemByTransactionID find all transaction item by transaction id
func (repo DatabaseRepository) FindAllTransactionItemByTransactionID(transactionID string) (transactionItems []models.TransactionItem, err error) {
	err = repo.db.Find(&transactionItems, "transaction_id = ?", transactionID).Error

	return
}

// UpdateTransactionItemByID update transaction item by id
func (repo DatabaseRepository) UpdateTransactionItemByID(id string, changeSet models.TransactionItemChangeSet) (err error) {
	var updates map[string]interface{}
	transactionItem := models.TransactionItem{
		Base: models.Base{ID: id},
	}

	mapstructure.Decode(changeSet, &updates)

	err = repo.db.Model(&transactionItem).Updates(updates).Error

	return
}

// DeleteTransactionItemByID delete transaction item by id
func (repo DatabaseRepository) DeleteTransactionItemByID(id string) (err error) {
	err = repo.db.Delete(&models.TransactionItem{}, "id = ?", id).Error

	return
}
