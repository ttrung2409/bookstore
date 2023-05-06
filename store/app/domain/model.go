package domain

import "github.com/google/uuid"

type DataObject interface{}

const EmptyId = ""

func NewId() string {
	return uuid.NewString()
}
