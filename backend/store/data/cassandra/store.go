package data

import (
	"store/data"

	"github.com/google/uuid"
)


var storeId = data.EntityId(uuid.NewString())

func StoreId() data.EntityId {
	return storeId
}