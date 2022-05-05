package data

import (
	"github.com/google/uuid"
)

type Model interface{}

const EmptyId = ""

func NewId() string {
	return uuid.NewString()
}
