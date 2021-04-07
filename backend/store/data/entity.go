package data

import (
	"github.com/google/uuid"
)


type EntityId string

func (id EntityId) ToString() string {
	return (string)(id)
}

func (id EntityId) ToMap() map[string]interface{} {
	return map[string]interface{}{"id": id}
}

const EmptyEntityId EntityId = ""

func NewEntityId() EntityId {
	return EntityId(uuid.NewString())
}