package models

// Transaction transaction struct
type Transaction struct {
	Base
	UserID string `json:"user_id"`
	Status int64  `json:"status"`
	Total  int64  `json:"total"`
}

// TransactionDetail Transaction detail struct
type TransactionDetail struct {
	Transaction
	Items []*TransactionItemDetail `json:"items"`
}

// TransactionCreateRequest transaction create request struct
type TransactionCreateRequest struct {
	UserID string `json:"user_id"`
}

// TransactionChangeSet transaction change set struct
type TransactionChangeSet struct {
	Status int64 `json:"status" mapstructure:",omitempty"`
	Total  int64 `json:"total" mapstructure:",omitempty"`
}

// NewTransaction new transaction from request
func NewTransaction(request TransactionCreateRequest) Transaction {
	status := int64(1)

	return Transaction{
		UserID: request.UserID,
		Status: status,
	}
}

// IsExistTransactionItemByItemID is exist transaction item by item id
func (transactionDetail TransactionDetail) IsExistTransactionItemByItemID(itemID string) bool {
	for _, item := range transactionDetail.Items {
		if item.ItemID == itemID {
			return true
		}
	}

	return false
}

// GetTransactionItemByItemID get transaction item by item id
func (transactionDetail TransactionDetail) GetTransactionItemByItemID(itemID string) *TransactionItem {
	for _, item := range transactionDetail.Items {
		if item.ItemID == itemID {
			return &item.TransactionItem
		}
	}

	return nil
}

// IndexTransactionItemByItemID index transaction item by item id
func (transactionDetail TransactionDetail) IndexTransactionItemByItemID(itemID string) int64 {
	for i, item := range transactionDetail.Items {
		if item.ItemID == itemID {
			return int64(i)
		}
	}

	return -1
}
