package repository

import "github.com/google/uuid"

var storeId = uuid.NewString()

func StoreId() string {
	return storeId
}