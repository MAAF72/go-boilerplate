package repositories

import (
	"github.com/MAAF72/go-boilerplate/models"
	"github.com/mitchellh/mapstructure"
)

// UserRepository user repository interface
type UserRepository interface {
	SaveUser(request models.User) (id string, err error)
	FindUserByID(id string) (user models.User, err error)
	FindUserByPhoneNumber(phoneNumber string) (user models.User, err error)
	UpdateUserByID(id string, changeSet models.UserChangeSet) (err error)
	DeleteUserByID(id string) (err error)
}

// SaveUser save user
func (repo DatabaseRepository) SaveUser(user models.User) (id string, err error) {
	err = repo.db.Create(&user).Error

	id = user.ID

	return
}

// FindUserByID find user by id
func (repo DatabaseRepository) FindUserByID(id string) (user models.User, err error) {
	err = repo.db.Take(&user, "id = ?", id).Error

	return
}

// FindUserByPhoneNumber find user by phone number
func (repo DatabaseRepository) FindUserByPhoneNumber(phoneNumber string) (user models.User, err error) {
	err = repo.db.Take(&user, "phone_number = ?", phoneNumber).Error

	return
}

// UpdateUserByID update user by id
func (repo DatabaseRepository) UpdateUserByID(id string, changeSet models.UserChangeSet) (err error) {
	var updates map[string]interface{}
	user := models.User{
		Base: models.Base{ID: id},
	}

	mapstructure.Decode(changeSet, &updates)

	err = repo.db.Model(&user).Updates(updates).Error

	return
}

// DeleteUserByID delete user by id
func (repo DatabaseRepository) DeleteUserByID(id string) (err error) {
	err = repo.db.Delete(&models.User{}, "id = ?", id).Error

	return
}
