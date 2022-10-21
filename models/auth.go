package models

import "golang.org/x/crypto/bcrypt"

// UserRegisterRequest user register request struct
type UserRegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password" mapstructure:",omitempty"`
}

// UserRegisterResponse user register response struct
type UserRegisterResponse struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}

// UserLoginRequest user login request struct
type UserLoginRequest struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Password    string `json:"password"`
}

// NewUserFromRegister new user from register request
func NewUserFromRegister(request UserRegisterRequest) User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	return User{
		PhoneNumber: request.PhoneNumber,
		Name:        request.Name,
		Role:        request.Role,
		Hash:        string(hash),
	}
}
