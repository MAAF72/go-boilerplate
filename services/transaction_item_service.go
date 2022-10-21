package services

import (
	"context"

	"github.com/MAAF72/go-boilerplate/models"
)

// TransactionItemService transaction item service interface
type TransactionItemService interface {
	CreateTransactionItem(ctx context.Context, request models.TransactionItemCreateRequest) (res models.TransactionItem, err error)
	GetTransactionItemByID(ctx context.Context, id string) (res models.TransactionItem, err error)
	GetTransactionItemDetailByID(ctx context.Context, id string) (res models.TransactionItemDetail, err error)
	GetAllTransactionItemByTransactionID(ctx context.Context, transactionID string) (res []models.TransactionItem, err error)
	GetAllTransactionItemDetailByTransactionID(ctx context.Context, transactionID string) (res []*models.TransactionItemDetail, err error)
	UpdateTransactionItemByID(ctx context.Context, id string, changeSet models.TransactionItemChangeSet) (res models.TransactionItem, err error)
	DeleteTransactionItemByID(ctx context.Context, id string) (err error)
}

// CreateTransactionItem create transaction item
func (service Service) CreateTransactionItem(ctx context.Context, request models.TransactionItemCreateRequest) (res models.TransactionItem, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	newTransactionItem := models.NewTransactionItem(request)

	item, err := service.GetItemByID(ctx, newTransactionItem.ItemID)
	if err != nil {
		return
	}

	newTransactionItem.Price = item.Price

	id, err := dbRepo.SaveTransactionItem(newTransactionItem)
	if err != nil {
		return
	}

	res, err = service.GetTransactionItemByID(ctx, id)

	return
}

// GetTransactionItemByID get transaction item by id
func (service Service) GetTransactionItemByID(ctx context.Context, id string) (res models.TransactionItem, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.FindTransactionItemByID(id)
}

// GetTransactionItemDetailByID get transaction item detail by id
func (service Service) GetTransactionItemDetailByID(ctx context.Context, id string) (res models.TransactionItemDetail, err error) {
	transactionItem, err := service.GetTransactionItemByID(ctx, id)
	if err != nil {
		return
	}

	transaction, err := service.GetTransactionByID(ctx, transactionItem.TransactionID)
	if err != nil {
		return
	}

	item, err := service.GetItemByID(ctx, transactionItem.ItemID)
	if err != nil {
		return
	}

	res = models.TransactionItemDetail{
		TransactionItem: transactionItem,
		Transaction:     &transaction,
		Item:            &item,
	}

	return
}

// GetAllTransactionItemByTransactionID get all transaction item by transaction id
func (service Service) GetAllTransactionItemByTransactionID(ctx context.Context, transactionID string) (res []models.TransactionItem, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.FindAllTransactionItemByTransactionID(transactionID)
}

// GetAllTransactionItemDetailByTransactionID get all transaction item detail by transaction id
func (service Service) GetAllTransactionItemDetailByTransactionID(ctx context.Context, transactionID string) (res []*models.TransactionItemDetail, err error) {
	transactionItems, err := service.GetAllTransactionItemByTransactionID(ctx, transactionID)
	if err != nil {
		return
	}

	for _, transactionItem := range transactionItems {
		transactionItemDetail, err := service.GetTransactionItemDetailByID(ctx, transactionItem.ID)
		if err != nil {
			return []*models.TransactionItemDetail{}, err
		}

		res = append(res, &transactionItemDetail)
	}

	return
}

// UpdateTransactionItemByID update transaction item by id
func (service Service) UpdateTransactionItemByID(ctx context.Context, id string, changeSet models.TransactionItemChangeSet) (transactionItem models.TransactionItem, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	err = dbRepo.UpdateTransactionItemByID(id, changeSet)
	if err != nil {
		return
	}

	return dbRepo.FindTransactionItemByID(id)
}

// DeleteTransactionItemByID delete transaction item by id
func (service Service) DeleteTransactionItemByID(ctx context.Context, id string) (err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.DeleteTransactionItemByID(id)
}
