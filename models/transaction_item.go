package models

// TransactionItem transaction item struct
type TransactionItem struct {
	Base
	TransactionID  string `json:"transaction_id"`
	ItemID         string `json:"item_id"`
	Quantity       int64  `json:"quantity"`
	Price          int64  `json:"price"`
	Discount       int64  `json:"discount"`
	DiscountReason string `json:"discount_reason"`
}

// TransactionItemDetail transaction item detail struct
type TransactionItemDetail struct {
	TransactionItem
	Transaction *Transaction `json:"transaction"`
	Item        *Item        `json:"item"`
}

// TransactionItemCreateRequest transaction item create request struct
type TransactionItemCreateRequest struct {
	TransactionID  string `json:"transaction_id"`
	ItemID         string `json:"item_id"`
	Quantity       int64  `json:"quantity"`
	Price          int64  `json:"price"`
	Discount       int64  `json:"discount"`
	DiscountReason string `json:"discount_reason"`
}

// TransactionItemChangeSet transaction item change set struct
type TransactionItemChangeSet struct {
	Quantity       int64  `json:"quantity" mapstructure:",omitempty"`
	Price          int64  `json:"price" mapstructure:",omitempty"`
	Discount       int64  `json:"discount" mapstructure:",omitempty"`
	DiscountReason string `json:"discount_reason" mapstructure:",omitempty"`
}

// NewTransactionItem new transaction item from request
func NewTransactionItem(request TransactionItemCreateRequest) TransactionItem {
	return TransactionItem{
		TransactionID:  request.TransactionID,
		ItemID:         request.ItemID,
		Quantity:       request.Quantity,
		Price:          request.Price,
		Discount:       request.Discount,
		DiscountReason: request.DiscountReason,
	}
}
