package models

// User user struct
type User struct {
	Base
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Hash        string `json:"-"`
}

// UserChangeSet user change set struct
type UserChangeSet struct {
	Name     string `json:"name" mapstructure:",omitempty"`
	Password string `json:"password" mapstructure:",omitempty"`
}
