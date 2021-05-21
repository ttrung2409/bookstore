package domain

import (
	"github.com/google/uuid"
)

type Entity interface {
	GetId() EntityId
	SetId(id EntityId)
}

type EntityId string

func (id EntityId) ToString() string {
	return string(id)
}

const EmptyEntityId EntityId = ""

func NewEntityId() EntityId {
	return EntityId(uuid.NewString())
}

func FromStringToEntityId(id string) EntityId {
	return EntityId(id)
}
