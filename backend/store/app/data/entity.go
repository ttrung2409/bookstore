package data

import (
	"github.com/google/uuid"
)

type Identifier interface {
	Value() EntityId
	ToMap() map[string]interface{}
}

type EntityId string

func (id EntityId) Value() EntityId {
	return (EntityId)(id)
}

func (id EntityId) ToMap() map[string]interface{} {
	return map[string]interface{}{"id": id}
}

func (id EntityId) ToString() string {
	return (string)(id)
}

const EmptyEntityId EntityId = ""

func NewEntityId() EntityId {
	return EntityId(uuid.NewString())
}
