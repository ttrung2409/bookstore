package operation

import (
	"time"
)

type Order struct {
	Id         string
	Number     string
	CreatedAt  time.Time
	CustomerId string
	Status     int
}
