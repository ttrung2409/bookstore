package domain

import (
	"github.com/google/uuid"
)

type DataObject any

func NewId() string {
	return uuid.NewString()
}
