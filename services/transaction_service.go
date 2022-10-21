package services

import (
	"context"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/repositories"
)

// TransactionService transaction service interface
type TransactionService interface {
	CreateTransaction(ctx context.Context, request models.TransactionCreateRequest) (res models.Transaction, err error)
	GetTransactionByID(ctx context.Context, id string) (res models.Transaction, err error)
	GetTransactionDetailByID(ctx context.Context, id string) (res models.TransactionDetail, err error)
	UpdateTransactionByID(ctx context.Context, id string, changeSet models.TransactionChangeSet) (res models.Transaction, err error)
	DeleteTransactionByID(ctx context.Context, id string) (err error)
	TransactionApplyPromo(ctx context.Context, id string) (res models.TransactionDetail, err error)
}

// CreateTransaction create transaction
func (service Service) CreateTransaction(ctx context.Context, request models.TransactionCreateRequest) (res models.Transaction, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	newTransaction := models.NewTransaction(request)

	id, err := dbRepo.SaveTransaction(newTransaction)
	if err != nil {
		return
	}

	res, err = service.GetTransactionByID(ctx, id)

	return
}

// GetTransactionByID get transaction by id
func (service Service) GetTransactionByID(ctx context.Context, id string) (res models.Transaction, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.FindTransactionByID(id)
}

// GetTransactionDetailByID get transaction detail by id
func (service Service) GetTransactionDetailByID(ctx context.Context, id string) (res models.TransactionDetail, err error) {
	transaction, err := service.GetTransactionByID(ctx, id)
	if err != nil {
		return
	}

	transactionItemDetails, err := service.GetAllTransactionItemDetailByTransactionID(ctx, transaction.ID)
	if err != nil {
		return
	}

	res = models.TransactionDetail{
		Transaction: transaction,
		Items:       transactionItemDetails,
	}

	return
}

// UpdateTransactionByID update transaction by id
func (service Service) UpdateTransactionByID(ctx context.Context, id string, changeSet models.TransactionChangeSet) (transaction models.Transaction, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	err = dbRepo.UpdateTransactionByID(id, changeSet)
	if err != nil {
		return
	}

	return dbRepo.FindTransactionByID(id)
}

// DeleteTransactionByID delete transaction by id
func (service Service) DeleteTransactionByID(ctx context.Context, id string) (err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.DeleteTransactionByID(id)
}

// TransactionApplyPromo transaction apply promo
func (service Service) TransactionApplyPromo(ctx context.Context, id string) (res models.TransactionDetail, err error) {
	dbRepo := DatabaseRepositories(ctx, service)
	gruleRepo := service.repositories.GruleRepository()

	gruleKnowledgeLib := gruleRepo.NewGruleKnowledgeBaseInstance("discount", "0.0.1")
	gruleEngine := gruleRepo.NewGruleEngine()
	gruleDataContext := gruleRepo.NewGruleDataContext()

	// read table => transaction_item (id)
	transactionDetails, err := service.GetTransactionDetailByID(ctx, id)
	if err != nil {
		return
	}

	// discountnya di 0 kan dulu semua

	gruleDataContext.Add("transaction", &transactionDetails)
	gruleDataContext.AddJSON("temp", []byte("{}"))
	err = gruleEngine.Execute(gruleDataContext, gruleKnowledgeLib)
	if err != nil {
		return
	}

	err = dbRepo.Transaction(func(modelRepo repositories.DatabaseRepositoriesImpl) error {
		// update table => transaction (total), transaction_item (discount, discount_reason)

		return nil
	})

	res = transactionDetails

	return
}
