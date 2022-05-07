package integration

import (
	"encoding/json"
	"store/kafka"
)

type message struct{}

func (m message) Value() []byte {
	serialized, _ := json.Marshal(m)

	return serialized
}

func Deserialize[M kafka.Message](msg kafka.Message, out M) M {
	json.Unmarshal(msg.Value(), out)

	return out
}
