package messages

import (
	"encoding/json"
)

type message struct{}

func (m *message) Value() []byte {
	serialized, _ := json.Marshal(m)

	return serialized
}
