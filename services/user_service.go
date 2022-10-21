package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/MAAF72/go-boilerplate/models"
	"github.com/MAAF72/go-boilerplate/repositories"
	gonanoid "github.com/matoous/go-nanoid"
)

// UserService user service interface
type UserService interface {
	RegisterUser(ctx context.Context, request models.UserRegisterRequest) (res models.UserRegisterResponse, err error)
	GetUserByID(ctx context.Context, id string) (res models.User, err error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (res models.User, err error)
	UpdateUserByID(ctx context.Context, id string, changeSet models.UserChangeSet) (user models.User, err error)
	UpdateUserByID2(ctx context.Context, id string, changeSet models.UserChangeSet) (user models.User, err error)
	DeleteUserByID(ctx context.Context, id string) (err error)
}

// RegisterUser register
func (service Service) RegisterUser(ctx context.Context, request models.UserRegisterRequest) (res models.UserRegisterResponse, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	// find another user with same phone number
	userWithSamePhoneNumber, _ := dbRepo.FindUserByPhoneNumber(request.PhoneNumber)

	if userWithSamePhoneNumber.Base.ID != "" {
		err = errors.New("phone number already registered")
		return
	}

	// generate random password if password is not provided
	if request.Password == "" {
		// generate random password (4 characters)
		generatedPassword, _ := gonanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyz", 4)

		request.Password = generatedPassword
	}

	newUser := models.NewUserFromRegister(request)

	err = dbRepo.Transaction(func(modelRepo repositories.DatabaseRepositoriesImpl) error {
		// if you want to make it single transaction, keep the repo from the previous ctx if exist
		// if you want to make it nested transaction, update the repo to the new one from parameter and ignore the error result of nested transaction
		// if ctxCurrRepo := ctx.Value(DBREPOKEY); ctxCurrRepo != nil {
		// 	if currRepo, ok := ctxCurrRepo.(repositories.DatabaseRepositoriesImpl); ok {
		// 		modelRepo = currRepo
		// 	}
		// }

		newCtx := context.WithValue(ctx, DBREPOKEY, modelRepo)
		newCtx = context.WithValue(newCtx, "src", "CreateUser")

		id, err := modelRepo.SaveUser(newUser)
		if err != nil {
			return err
		}

		user, err := service.UpdateUserByID(newCtx, id, models.UserChangeSet{
			Name: fmt.Sprintf("%s-oke", newUser.Name),
		})
		if err != nil {
			return err
		}

		fmt.Println("Panggil service.UpdateUserByID2")
		user, err = service.UpdateUserByID2(newCtx, id, models.UserChangeSet{
			Name: fmt.Sprintf("%s-oke2", user.Name),
		})
		// comment the err check below if you want the update above is ignored but the main transaction is commited
		// if err != nil {
		// 	return err
		// }

		fmt.Println("Panggil service.UpdateUserByID2")
		user, err = service.UpdateUserByID2(newCtx, id, models.UserChangeSet{
			Name: fmt.Sprintf("%s-oke3", user.Name),
		})
		if err != nil {
			return err
		}

		user, err = service.GetUserByID(newCtx, id)

		res = models.UserRegisterResponse{
			Base:        user.Base,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
			Role:        user.Role,
			Password:    request.Password,
		}

		// err = errors.New("iseng ja")

		return err
	})
	if err != nil {
		return
	}

	return
}

// GetUserByID get user by id
func (service Service) GetUserByID(ctx context.Context, id string) (res models.User, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	fmt.Println("src", ctx.Value("src"))
	fmt.Println("repo ctx", ctx.Value(DBREPOKEY))
	fmt.Println("repo", &dbRepo)

	return dbRepo.FindUserByID(id)
}

// GetUserByPhoneNumber get user by phone number
func (service Service) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (res models.User, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.FindUserByPhoneNumber(phoneNumber)
}

// UpdateUserByID update user by id
func (service Service) UpdateUserByID(ctx context.Context, id string, changeSet models.UserChangeSet) (res models.User, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	fmt.Println("Hadir UpdateUserByID")

	err = dbRepo.UpdateUserByID(id, changeSet)
	if err != nil {
		return
	}

	return dbRepo.FindUserByID(id)
}

// UpdateUserByID2 update user by id
func (service Service) UpdateUserByID2(ctx context.Context, id string, changeSet models.UserChangeSet) (res models.User, err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	err = dbRepo.Transaction(func(modelRepo repositories.DatabaseRepositoriesImpl) error {
		// if you want to make it single transaction, keep the repo from the previous ctx if exist by uncomment the code below
		// if you want to make it nested transaction, update the repo to the new one from parameter by comment the code below
		// if ctxCurrRepo := ctx.Value(DBREPOKEY); ctxCurrRepo != nil {
		// 	if currRepo, ok := ctxCurrRepo.(repositories.DatabaseRepositoriesImpl); ok {
		// 		modelRepo = currRepo
		// 	}
		// }

		newCtx := context.WithValue(ctx, DBREPOKEY, modelRepo)
		newCtx = context.WithValue(newCtx, "src", "UpdateUserByID2")

		fmt.Println("Hadir UpdateUserByID2")

		err := modelRepo.UpdateUserByID(id, changeSet)
		if err != nil {
			return err
		}

		err = errors.New("hai petrik")

		return err
	})
	if err != nil {
		return
	}

	return dbRepo.FindUserByID(id)
}

// DeleteUserByID delete user by id
func (service Service) DeleteUserByID(ctx context.Context, id string) (err error) {
	dbRepo := DatabaseRepositories(ctx, service)

	return dbRepo.DeleteUserByID(id)
}
