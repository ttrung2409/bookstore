package data

import (
	"github.com/google/uuid"
)

type Entity interface {
	GetId() string
	SetId(id string)
}

const EmptyEntityId = ""

func NewEntityId() string {
	return uuid.NewString()
}
