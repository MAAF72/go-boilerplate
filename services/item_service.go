package services

import (
	"context"

	"github.com/MAAF72/go-boilerplate/models"
)

// ItemService item service interface
type ItemService interface {
	CreateItem(ctx context.Context, request models.ItemCreateRequest) (res models.Item, err error)
	GetItemByID(ctx context.Context, id string) (res models.Item, err error)
	UpdateItemByID(ctx context.Context, id string, changeSet models.ItemChangeSet) (res models.Item, err error)
	DeleteItemByID(ctx context.Context, id string) (err error)
}

// CreateItem create item
func (service Service) CreateItem(ctx context.Context, request models.ItemCreateRequest) (res models.Item, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	newItem := models.NewItem(request)

	id, err := dbRepo.SaveItem(newItem)
	if err != nil {
		return
	}

	res, err = service.GetItemByID(ctx, id)

	return
}

// GetItemByID get item by id
func (service Service) GetItemByID(ctx context.Context, id string) (res models.Item, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.FindItemByID(id)
}

// UpdateItemByID update item by id
func (service Service) UpdateItemByID(ctx context.Context, id string, changeSet models.ItemChangeSet) (item models.Item, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	err = dbRepo.UpdateItemByID(id, changeSet)
	if err != nil {
		return
	}

	return dbRepo.FindItemByID(id)
}

// DeleteItemByID delete item by id
func (service Service) DeleteItemByID(ctx context.Context, id string) (err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.DeleteItemByID(id)
}
