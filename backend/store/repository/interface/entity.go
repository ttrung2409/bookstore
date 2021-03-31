package repository

import "github.com/google/uuid"

type EntityId string

const EmptyEntityId EntityId = ""

func NewEntityId() EntityId {
	return EntityId(uuid.NewString())
}