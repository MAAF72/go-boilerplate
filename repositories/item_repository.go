package repositories

import (
	"github.com/MAAF72/go-boilerplate/models"
	"github.com/mitchellh/mapstructure"
)

// ItemRepository item repository interface
type ItemRepository interface {
	SaveItem(request models.Item) (id string, err error)
	FindItemByID(id string) (item models.Item, err error)
	UpdateItemByID(id string, changeSet models.ItemChangeSet) (err error)
	DeleteItemByID(id string) (err error)
}

// SaveItem save item
func (repo DatabaseRepository) SaveItem(item models.Item) (id string, err error) {
	err = repo.db.Create(&item).Error

	id = item.ID

	return
}

// FindItemByID find item by id
func (repo DatabaseRepository) FindItemByID(id string) (item models.Item, err error) {
	err = repo.db.Take(&item, "id = ?", id).Error

	return
}

// UpdateItemByID update item by id
func (repo DatabaseRepository) UpdateItemByID(id string, changeSet models.ItemChangeSet) (err error) {
	var updates map[string]interface{}
	item := models.Item{
		Base: models.Base{ID: id},
	}

	mapstructure.Decode(changeSet, &updates)

	err = repo.db.Model(&item).Updates(updates).Error

	return
}

// DeleteItemByID delete item by id
func (repo DatabaseRepository) DeleteItemByID(id string) (err error) {
	err = repo.db.Delete(&models.Item{}, "id = ?", id).Error

	return
}
