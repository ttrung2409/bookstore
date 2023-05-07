package domain

import "github.com/google/uuid"

type DataObject any

const EmptyId = ""

func NewId() string {
	return uuid.NewString()
}
