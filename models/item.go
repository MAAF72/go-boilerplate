package models

// Item item struct
type Item struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
}

// ItemCreateRequest item create request struct
type ItemCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
}

// ItemChangeSet item change set struct
type ItemChangeSet struct {
	Name        string `json:"name" mapstructure:",omitempty"`
	Description string `json:"description" mapstructure:",omitempty"`
	Quantity    int64  `json:"quantity" mapstructure:",omitempty"`
	Price       int64  `json:"price" mapstructure:",omitempty"`
}

// NewItem new item from request
func NewItem(request ItemCreateRequest) Item {
	return Item{
		Name:        request.Name,
		Description: request.Description,
		Quantity:    request.Quantity,
		Price:       request.Price,
	}
}
